package api

import (
	"go-rest-api/util"
	"net/http"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	util.RespondJSON(w, 200, map[string]string{"message": "Pong"})
}
