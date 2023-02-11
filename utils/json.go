package utils

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, status int, data map[string]interface{}) {
	bytes, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func Message(w http.ResponseWriter, status int, message string) {
	JSON(w, status, map[string]interface{}{
		"message": message,
	})
}

func Error(w http.ResponseWriter, status int, err error) {
	Message(w, status, err.Error())
}
