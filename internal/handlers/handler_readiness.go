package handler

import (
	"net/http"

	"github.com/psv2522/rss-aggregator/api"
)

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	api.RespondWithJSON(w, 200, struct{}{})
}
