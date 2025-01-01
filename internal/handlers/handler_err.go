package handler

import (
	"net/http"

	"github.com/psv2522/rss-aggregator/api"
)

func HandlerErr(w http.ResponseWriter, r *http.Request) {
	api.RespondWithError(w, 400, "Something went wrong")
}
