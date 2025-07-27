package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"
	"github.com/sahilrana7582/vitals-guard/gateway/middleware"
	"github.com/sahilrana7582/vitals-guard/gateway/routes"
)

func main() {
	mux := http.NewServeMux()

	routes.RegisterRoutes(mux)

	loggedMux := middleware.LoggingMiddleware(mux)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}).Handler(loggedMux)

	log.Println("🚀 Gateway running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}
