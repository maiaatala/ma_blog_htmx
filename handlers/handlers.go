package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"ssr-htmx/services"
	"ssr-htmx/views"

	"github.com/a-h/templ"
	// "github.com/a-h/templ"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	page := views.MainPage(r.URL.Path)

	err := page.Render(r.Context(), w)
	if err != nil {
		http.Error(w, "failed to render root", http.StatusInternalServerError)
	}
}

func AboutPartialHandler(w http.ResponseWriter, r *http.Request) {
	err := views.Layout(templ.NopComponent, views.About(), templ.NopComponent).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "failed to render partial", http.StatusInternalServerError)
	}
}

func ContactFormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "could not read body", http.StatusBadRequest)
		return
	}

	err = services.PostContactForm(body)
	if err != nil {
		http.Error(w, "failed to send", http.StatusBadGateway)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Contato enviado com sucesso!"))
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	err := views.NotFound().Render(r.Context(), w)
	if err != nil {
		http.Error(w, "failed to render 404 page", http.StatusInternalServerError)
	}
}

func InternalError(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	_ = views.InternalError().Render(r.Context(), w)
}

func WithNotFoundHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec := httptest.NewRecorder()
		next.ServeHTTP(rec, r)

		if rec.Code == http.StatusNotFound {
			NotFoundHandler(w, r)
			return
		}

		for k, v := range rec.Header() {
			w.Header()[k] = v
		}
		w.WriteHeader(rec.Code)
		_, _ = w.Write(rec.Body.Bytes())
	})
}
