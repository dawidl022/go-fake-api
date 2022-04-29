package server

import (
	"encoding/json"
	"net/http"
	"os"
	"server/models"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type data struct {
	a      []models.Album
	c      []models.Comment
	photos []models.Photo
	posts  []models.Post
	t      []models.Todo
	u      []models.User
}

type server struct {
	r chi.Router
	d *data
}

func StartServer() {
	s := newServer()
	s.r.Use(jsonHeaders)
	s.routes()

	http.ListenAndServe(":3000", s.r)
}

func loadData() *data {
	d := data{}

	rawAlbums, _ := os.ReadFile("data/albums.json")
	json.Unmarshal(rawAlbums, &d.a)

	rawComments, _ := os.ReadFile("data/comments.json")
	json.Unmarshal(rawComments, &d.c)

	rawPhotos, _ := os.ReadFile("data/photos.json")
	json.Unmarshal(rawPhotos, &d.photos)

	rawPosts, _ := os.ReadFile("data/posts.json")
	json.Unmarshal(rawPosts, &d.posts)

	rawTodos, _ := os.ReadFile("data/todos.json")
	json.Unmarshal(rawTodos, &d.t)

	rawUsers, _ := os.ReadFile("data/users.json")
	json.Unmarshal(rawUsers, &d.u)

	return &d
}

func newServer() server {
	s := server{
		r: chi.NewRouter(),
		d: loadData(),
	}

	s.r.Use(middleware.RequestID)
	s.r.Use(middleware.RealIP)
	s.r.Use(middleware.Logger)
	s.r.Use(middleware.Recoverer)

	s.r.Use(middleware.Timeout(60 * time.Second))

	return s
}

func jsonHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}
