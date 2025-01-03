package handler

import (
	"net/http"

	"github.com/psv2522/rss-aggregator/api"
)

func HandleErr(w http.ResponseWriter, r *http.Request) {
	api.RespondWithError(w, 400, "Something went wrong")
}
