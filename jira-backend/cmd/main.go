package main

import (
	"net/http"
	"encoding/json"
)

func handlerHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		response := map[string]string{
			"status" : "ok",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func main() {
	http.HandleFunc("/api/v1/health", handlerHealth)
	http.ListenAndServe(":8000", nil)
}