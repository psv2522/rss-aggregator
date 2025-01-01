package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/psv2522/rss-aggregator/api"
	"github.com/psv2522/rss-aggregator/internal/config"
	"github.com/psv2522/rss-aggregator/internal/database"
)

func HandlerCreateUser(apiCfg *config.ApiConfig, w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		api.RespondWithError(w, 400, "Error parsing JSON")
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        pgtype.UUID{Bytes: uuid.New(), Valid: true},
		CreatedAt: pgtype.Timestamp{Time: time.Now().UTC(), Valid: true},
		UpdatedAt: pgtype.Timestamp{Time: time.Now().UTC(), Valid: true},
		Name:      params.Name,
	})
	if err != nil {
		api.RespondWithError(w, 400, "Could not create user")
		return
	}

	api.RespondWithJSON(w, 201, user)
}

func MakeHandlerCreateUser(apiCfg config.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		HandlerCreateUser(&apiCfg, w, r)
	}
}
