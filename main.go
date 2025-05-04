package main

import (
	"log"
	"net/http"
	"ssr-htmx/views"

	"github.com/a-h/templ"
)

func main() {
	component := views.Index("Templ!")
	http.Handle("/", templ.Handler(component))

	// server static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Serving on port 3001")
	log.Fatal(http.ListenAndServe(":3001", nil))
}
