package server

import (
	"fmt"
	"net/http"
	"os"
)

func graphiql(basedir string) http.HandlerFunc {
	s, err := os.ReadFile(fmt.Sprintf("%sserver/templates/graphiql.html", basedir))

	return func(w http.ResponseWriter, r *http.Request) {
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
		}

		w.Write(s)
	}
}
