package handlers

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"ssrhtmx/services"
	"ssrhtmx/views"
	"strconv"
	"strings"
	"time"

	"github.com/a-h/templ"
	// "github.com/a-h/templ"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	page := views.MainPage(views.WithPath(r.URL.Path))

	err := page.Render(r.Context(), w)
	if err != nil {
		http.Error(w, "failed to render root", http.StatusInternalServerError)
	}
}

func AboutStatic(w http.ResponseWriter, r *http.Request) {
	about := views.Layout(templ.NopComponent, views.About(), templ.NopComponent)
	page := views.MainPage(views.WithChild(about))

	err := page.Render(r.Context(), w)
	if err != nil {
		http.Error(w, "failed to render about", http.StatusInternalServerError)
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

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	pageNum := 1

	data, _err := services.FetchPosts(pageNum)
	if _err != nil {
		http.Error(w, "Failed to load posts", http.StatusBadGateway)
		return
	}

	next := ""
	if len(data.Items) > 0 {
		next = "/partial/posts?page=" + fmt.Sprint(pageNum+1)
	}

	homePosts := views.PostList(data.Items, next)
	page := views.MainPage(views.WithPostChild(homePosts))

	err := page.Render(r.Context(), w)
	if err != nil {
		http.Error(w, "failed to render home page", http.StatusInternalServerError)
	}
}

func PartialHome(w http.ResponseWriter, r *http.Request) {
	err := views.PostsPartial().Render(r.Context(), w)
	if err != nil {
		http.Error(w, "failed to home content", http.StatusInternalServerError)
	}
}

func GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	pageNum := 1
	if val := r.URL.Query().Get("page"); val != "" {
		if n, err := strconv.Atoi(val); err == nil {
			pageNum = n
		}
	}

	data, err := services.FetchPosts(pageNum)
	if err != nil {
		http.Error(w, "Failed to load posts", http.StatusBadGateway)
		return
	}

	next := ""
	if len(data.Items) > 0 {
		next = "/partial/posts?page=" + fmt.Sprint(pageNum+1)
	}

	err = views.PostList(data.Items, next).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Failed to render posts", http.StatusInternalServerError)
	}
}

func PostPartialHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/partial/post/")
	post, err := services.FetchPostByID(id)
	parsed, err := time.Parse(time.RFC3339, post.UpdatedAt)
	if err == nil {
		post.UpdatedAt = parsed.Format("02/01/2006") // formato pt-BR: dd/mm/aaaa
	}

	if err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	err = views.PostDetailedPartial(*post).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Failed to render post partial", http.StatusInternalServerError)
	}
}

func PostPageHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/post/")
	post, err := services.FetchPostByID(id)
	parsed, err := time.Parse(time.RFC3339, post.UpdatedAt)
	if err == nil {
		post.UpdatedAt = parsed.Format("02/01/2006") // formato pt-BR: dd/mm/aaaa
	}
	if err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	content := views.PostDetailedPartial(*post)
	page := views.MainPage(views.WithChild(content))

	err = page.Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Failed to render post page", http.StatusInternalServerError)
	}
}

func PartialCommentsHandler(w http.ResponseWriter, r *http.Request) {
	postId := r.URL.Query().Get("postId")
	parentId := r.URL.Query().Get("parentCommentId") // pode estar vazio

	comments, err := services.FetchComments(postId, parentId)
	if err != nil {
		http.Error(w, "Failed to load comments", http.StatusInternalServerError)
		return
	}

	// nível +1 se for resposta
	level := 0
	if parentId != "" {
		level = 1
	}

	err = views.CommentList(comments, postId, level).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Failed to render comments", http.StatusInternalServerError)
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
