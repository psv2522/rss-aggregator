package handler

import (
	"net/http"

	"github.com/psv2522/rss-aggregator/api"
)

func HandleReadiness(w http.ResponseWriter, r *http.Request) {
	api.RespondWithJSON(w, 200, struct{}{})
}
