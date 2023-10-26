package web

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")

	response, err := json.MarshalIndent(payload, "", "    ")
	if err != nil {
		http.Error(w, "JSON Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(code)
	_, err = w.Write(response)
	if err != nil {
		log.Printf("Could not write json response: %s", err)
	}
}
