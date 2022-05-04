package server

import (
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func (s *server) routes(schema *graphql.Schema, basedir string) {
	s.router.Method(http.MethodPost, "/query", &relay.Handler{Schema: schema})
	s.router.Get("/graphql", graphiql(basedir))
}
