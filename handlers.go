package main

import "net/http"

// handlerReadiness handles the readiness check endpoint.
// It responds with a JSON object indicating the service status.
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

// handlerErr handles the error endpoint.
// It responds with a JSON object indicating an internal server error.
func handlerErr(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}
