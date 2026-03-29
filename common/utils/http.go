package util

import (
	"encoding/json"
	"maps"
	"net/http"
)

func ReadJSON[Body any](w http.ResponseWriter, r *http.Request) (Body, error) {
	const payloadLimit = 1 << 20
	body := http.MaxBytesReader(w, r.Body, payloadLimit)

	var v Body
	err := json.NewDecoder(body).Decode(&v)
	return v, err
}

func WriteJSON(w http.ResponseWriter, status int, v any, header ...http.Header) error {
	for _, h := range header {
		maps.Copy(w.Header(), h)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
