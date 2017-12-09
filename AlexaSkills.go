package main

import (
	"net/http"
	"os"

	"AlexaSkills/mux"
)

// Crowdbotics 
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	server := http.Server{
		Addr:    ":" + port,
		Handler: mux.NewHandler(),
	}

	server.ListenAndServe()
}

