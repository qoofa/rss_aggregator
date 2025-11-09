package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func responseWriteError(w http.ResponseWriter, statuscode int, msg string) {
	if statuscode > 499 {
		log.Println("Responding with 5XX error", msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, statuscode, errResponse{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, statuscode int, payload any) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshall JSONresponse: %v\n", payload)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statuscode)
	w.Write(data)
}
