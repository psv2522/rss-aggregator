package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/psv2522/rss-aggregator/internal/config"
	"github.com/psv2522/rss-aggregator/internal/database"
	handler "github.com/psv2522/rss-aggregator/internal/handlers"
)

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found in environment")
	}

	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		log.Fatal("DATABASE_URL is not found in environment")
	}

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dbUrl)
	if err != nil {
		log.Fatal("Could not connect to database:", err)
	}
	defer conn.Close(ctx)

	apiCfg := config.ApiConfig{
		DB: database.New(conn),
	}

	router := chi.NewRouter()
	router.Use(
		cors.Handler(
			cors.Options{
				AllowedOrigins:   []string{"https://*", "http://*"},
				AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
				AllowedHeaders:   []string{"*"},
				ExposedHeaders:   []string{"Link"},
				AllowCredentials: false,
				MaxAge:           300,
			},
		),
	)

	v1router := chi.NewRouter()
	v1router.Get("/healthz", handler.HandleReadiness)
	v1router.Get("/err", handler.HandleErr)
	v1router.Post("/users", handler.HandleCreateUser(apiCfg))
	v1router.Get("/users", handler.HandleGetUser(apiCfg))
	router.Mount("/v1", v1router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("Server starting on port %v", port)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}