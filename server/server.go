package server

import (
	"encoding/json"
	"fmt"
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
	s := newServer("")
	s.r.Use(jsonHeaders)
	s.routes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	http.ListenAndServe(":"+port, s.r)
}

func loadData(baseDir string) *data {
	d := data{}
	dir := baseDir + "data/"

	rawAlbums, _ := os.ReadFile(fmt.Sprintf("%salbums.json", dir))
	json.Unmarshal(rawAlbums, &d.a)

	rawComments, _ := os.ReadFile(fmt.Sprintf("%scomments.json", dir))
	json.Unmarshal(rawComments, &d.c)

	rawPhotos, _ := os.ReadFile(fmt.Sprintf("%sphotos.json", dir))
	json.Unmarshal(rawPhotos, &d.photos)

	rawPosts, _ := os.ReadFile(fmt.Sprintf("%sposts.json", dir))
	json.Unmarshal(rawPosts, &d.posts)

	rawTodos, _ := os.ReadFile(fmt.Sprintf("%stodos.json", dir))
	json.Unmarshal(rawTodos, &d.t)

	rawUsers, _ := os.ReadFile(fmt.Sprintf("%susers.json", dir))
	json.Unmarshal(rawUsers, &d.u)

	return &d
}

func newServer(baseDir string) server {
	s := server{
		r: chi.NewRouter(),
		d: loadData(baseDir),
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
