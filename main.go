package main

import (
	"log"
	"net/http"
	"os"
	"ssr-htmx/views"

	"github.com/a-h/templ"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	component := views.Index("Templ!")
	http.Handle("/", templ.Handler(component))

	// server static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	port := os.Getenv("PORT")
	log.Println("Serving on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
