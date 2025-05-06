package routes

import (
	"net/http"
	"ssr-htmx/handlers"
)

func SetupRoutes(mux *http.ServeMux) {
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/about", handlers.RootHandler)
	mux.HandleFunc("/about2", handlers.AboutStatic)
	mux.HandleFunc("/partial/about", handlers.AboutPartialHandler)
	mux.HandleFunc("/api/contact", handlers.ContactFormHandler)
}
