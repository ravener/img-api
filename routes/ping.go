package routes

import (
	"net/http"

	"github.com/ravener/img-api/utils"
)

// Ping is used to check if the server is alive.
func Ping(w http.ResponseWriter, r *http.Request) {
	utils.Message(w, http.StatusOK, "Pong!")
}
