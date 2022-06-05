package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"server/config"
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

func StartServer(conf *config.Config) {
	s := newServer()
	s.setup(conf)

	log.Fatal(http.ListenAndServe(":8080", s.router))
}

func (s *server) setup(conf *config.Config) {
	b, err := concatFiles(fmt.Sprintf("%sserver/graphql", conf.BaseDir),
		"query.graphql", "album.graphql", "post.graphql")
	if err != nil {
		log.Fatal("Cannot read grapql schema files:", err)
	}

	db, err := initDB(conf)
	if err != nil {
		log.Fatal("Cannot initialise database:", err)
	}

	schema := graphql.MustParseSchema(string(b), resolvers.NewRootResolver(db))
	s.routes(schema, conf.BaseDir)
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
