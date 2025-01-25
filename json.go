package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errorResponse{
		Error: msg,
	})
}

// respondWithJSON sends a JSON response with the specified HTTP status code and payload.
// It sets the "Content-Type" header to "application/json" and marshals the payload into JSON format.
// If marshalling fails, it logs the error and sends a 500 Internal Server Error response.
//
// Parameters:
//   - w: The http.ResponseWriter to write the response to.
//   - code: The HTTP status code to send.
//   - payload: The data to be marshalled into JSON and sent in the response body.
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	data, error := json.Marshal(payload)
	if error != nil {
		log.Printf("Error marshalling JSON: %s", error)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	w.Write(data)
}
