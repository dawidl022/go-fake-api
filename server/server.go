package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"server/resolvers"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	graphql "github.com/graph-gophers/graphql-go"
)

type server struct {
	router chi.Router
}

func newServer() server {
	s := server{
		router: chi.NewRouter(),
	}

	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.RealIP)
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)

	s.router.Use(middleware.Timeout(60 * time.Second))

	return s
}

func StartServer() {
	b, err := concatFiles("server/graphql", "query.graphql", "album.graphql", "post.graphql")
	if err != nil {
		log.Fatal("Cannot read grapql schema files")
	}

	root := resolvers.NewRootResolver()
	schema := graphql.MustParseSchema(string(b), root)

	s := newServer()
	s.routes(schema)

	log.Fatal(http.ListenAndServe(":8080", s.router))
}

func concatFiles(dirname string, filenames ...string) (string, error) {
	var res []byte

	for _, filename := range filenames {
		b, err := os.ReadFile(fmt.Sprintf("%s/%s", dirname, filename))
		if err != nil {
			return string(res), err
		}
		res = append(res, b...)
	}

	return string(res), nil
}
