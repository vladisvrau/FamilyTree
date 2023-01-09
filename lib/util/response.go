package util

import (
	"encoding/json"

	"net/http"
)

func WriteJson(w http.ResponseWriter, status int, content interface{}) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	err := json.NewEncoder(w).Encode(content)
	if err != nil {
		println(err.Error())
	}
}
