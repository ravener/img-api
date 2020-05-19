package routes

import (
	"net/http"
)

// Ping is used to check if the server is alive.
func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("{\"message\": \"Pong!\"}"))
}
