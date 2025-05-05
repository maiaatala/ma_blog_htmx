package handlers

import (
	"net/http"
	"net/http/httptest"
	"ssr-htmx/views"
	// "github.com/a-h/templ"
)

func AboutPartialHandler(w http.ResponseWriter, r *http.Request) {
	err := views.About().Render(r.Context(), w)
	if err != nil {
		http.Error(w, "failed to render partial", http.StatusInternalServerError)
	}
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	page := views.MainPage(r.URL.Path)

	err := page.Render(r.Context(), w)
	if err != nil {
		http.Error(w, "failed to render root", http.StatusInternalServerError)
	}
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
