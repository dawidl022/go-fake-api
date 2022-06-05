package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dawidl022/go-fake-api/config"
)

var testConf = &config.Config{
	BaseDir:     "../",
	DatabaseUrl: "postgresql://user:password@localhost:5432/fake?sslmode=disable",
}

func TestHandleGraphiql(t *testing.T) {
	srv := newServer()
	srv.setup(testConf)

	req := httptest.NewRequest("GET", "/graphql", nil)
	w := httptest.NewRecorder()
	srv.router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestHandleQuery(t *testing.T) {
	srv := newServer()
	srv.setup(testConf)

	req := httptest.NewRequest("POST", "/query", strings.NewReader(
		`{"operationName":null,"variables":{},"query":"{\n  album(id: 1) {\n    title\n    id\n    userId\n  }\n}\n"}`))
	req.Header.Set("Content-type", "application/json")
	w := httptest.NewRecorder()

	srv.router.ServeHTTP(w, req)
	assert.Equal(t, `{"data":{"album":{"title":"quidem molestiae enim","id":"1","userId":"1"}}}`,
		w.Body.String())
}
