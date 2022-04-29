package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleAlbums(t *testing.T) {
	srv := newServer("../")
	srv.routes()

	req := httptest.NewRequest("GET", "/albums", nil)
	w := httptest.NewRecorder()
	srv.r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestHandleUserById(t *testing.T) {
	srv := newServer("../")
	srv.routes()

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
	srv.r.ServeHTTP(w, req)

	assert.Equal(t, []byte(`{"Id":1,"Name":"Leanne Graham","Username":"Bret","Email":"Sincere@april.biz","Address":{"Street":"Kulas Light","Suite":"Apt. 556","City":"Gwenborough","Zipcode":"92998-3874","Geo":{"Lat":"-37.3159","Lng":"81.1496"}},"Phone":"1-770-736-8031 x56442","Website":"hildegard.org","Company":{"Name":"Romaguera-Crona","CatchPhrase":"Multi-layered client-server neural-net","Bs":"harness real-time e-markets"}}`),
		w.Body.Bytes())
}
