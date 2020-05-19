package utils

import (
	"net/http"
	"encoding/json"
)

func JSON(w http.ResponseWriter, status int, data map[string]interface{}) {
	bytes, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	w.WriteHeader(status)
	w.Write(bytes)
}
