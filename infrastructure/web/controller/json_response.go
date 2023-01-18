package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Fatal(err)
	}
}
