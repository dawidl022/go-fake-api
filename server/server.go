package server

import (
	"log"
	"net/http"
	"os"
	"server/resolvers"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func graphiql(w http.ResponseWriter, r *http.Request) {
	s, _ := os.ReadFile("server/templates/graphiql.html")
	w.Write(s)
}

func StartServer() {
	b, err := os.ReadFile("server/graphql/query.graphql")
	if err != nil {
		log.Fatal("Cannot read grapql schema file")
	}

	schema := graphql.MustParseSchema(string(b), &resolvers.Query{})
	http.Handle("/query", &relay.Handler{Schema: schema})
	http.HandleFunc("/graphql", graphiql)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"os"
// 	"server/models"
// 	"time"

// 	"github.com/go-chi/chi/v5"
// 	"github.com/go-chi/chi/v5/middleware"
// )

// type data struct {
// 	a      []models.Album
// 	c      []models.Comment
// 	photos []models.Photo
// 	posts  []models.Post
// 	t      []models.Todo
// 	u      []models.User
// }

// type server struct {
// 	r chi.Router
// 	d *data
// }

// func StartServer() {
// 	s := newServer("")
// 	s.r.Use(jsonHeaders)
// 	s.routes()

// 	http.ListenAndServe(":3000", s.r)
// }

// func loadData(baseDir string) *data {
// 	d := data{}
// 	dir := baseDir + "data/"

// 	rawAlbums, _ := os.ReadFile(fmt.Sprintf("%salbums.json", dir))
// 	json.Unmarshal(rawAlbums, &d.a)

// 	rawComments, _ := os.ReadFile(fmt.Sprintf("%scomments.json", dir))
// 	json.Unmarshal(rawComments, &d.c)

// 	rawPhotos, _ := os.ReadFile(fmt.Sprintf("%sphotos.json", dir))
// 	json.Unmarshal(rawPhotos, &d.photos)

// 	rawPosts, _ := os.ReadFile(fmt.Sprintf("%sposts.json", dir))
// 	json.Unmarshal(rawPosts, &d.posts)

// 	rawTodos, _ := os.ReadFile(fmt.Sprintf("%stodos.json", dir))
// 	json.Unmarshal(rawTodos, &d.t)

// 	rawUsers, _ := os.ReadFile(fmt.Sprintf("%susers.json", dir))
// 	json.Unmarshal(rawUsers, &d.u)

// 	return &d
// }

// func newServer(baseDir string) server {
// 	s := server{
// 		r: chi.NewRouter(),
// 		d: loadData(baseDir),
// 	}

// 	s.r.Use(middleware.RequestID)
// 	s.r.Use(middleware.RealIP)
// 	s.r.Use(middleware.Logger)
// 	s.r.Use(middleware.Recoverer)

// 	s.r.Use(middleware.Timeout(60 * time.Second))

// 	return s
// }

// func jsonHeaders(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/json; charset=utf-8")
// 		next.ServeHTTP(w, r)
// 	})
// }
