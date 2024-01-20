package handlers

import (
	"encoding/json"
	"io"
)

func ReadJSON(i interface{}, r io.Reader) error {
	return json.NewDecoder(r).Decode(i)
}

func WriteJSON(i interface{}, w io.Writer) error {
	return json.NewEncoder(w).Encode(i)
}
