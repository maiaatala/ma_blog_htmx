package main

import (
	"log"
	"net/http"
	"os"
	"ssrhtmx/handlers"
	"ssrhtmx/routes"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	// routes.SetupRoutes()

	// fs := http.FileServer(http.Dir("./static"))
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}
	log.Println("Serving on port " + port)

	mux := http.NewServeMux()
	routes.SetupRoutes(mux)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: handlers.WithNotFoundHandler(mux),
	}

	log.Fatal(server.ListenAndServe())

}
