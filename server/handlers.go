package server

import (
	"net/http"
	"os"
)

func graphiql(w http.ResponseWriter, r *http.Request) {
	s, err := os.ReadFile("server/templates/graphiql.html")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}

	w.Write(s)
}
