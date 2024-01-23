package handlers

import (
	"encoding/json"
	"io"
	"net/http"
)

func readJSON(r io.Reader, i interface{}) error {
	return json.NewDecoder(r).Decode(i)
}

func respondWithJSON(w http.ResponseWriter, status int, i interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(i)
}

func respondWithError(w http.ResponseWriter, status int, msg string) {
	type errorMsg struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, status, errorMsg{Error: msg})
}
