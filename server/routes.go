package server

import (
	"github.com/go-chi/chi/v5"
)

func (s *server) routes() {
	s.r.Route("/albums", func(r chi.Router) {
		r.Get("/", getAll(s.d.a))
		r.Get("/{id}", getById(s.d.a))
	})

	s.r.Route("/comments", func(r chi.Router) {
		r.Get("/", getAll(s.d.c))
		r.Get("/{id}", getById(s.d.c))
	})

	s.r.Route("/photos", func(r chi.Router) {
		r.Get("/", getAll(s.d.photos))
		r.Get("/{id}", getById(s.d.photos))
	})

	s.r.Route("/posts", func(r chi.Router) {
		r.Get("/", getAll(s.d.posts))
		r.Get("/{id}", getById(s.d.posts))
		r.Get("/{id}/comments", s.getCommentsByPostId())
	})

	s.r.Route("/todos", func(r chi.Router) {
		r.Get("/", getAll(s.d.t))
		r.Get("/{id}", getById(s.d.t))
	})

	s.r.Route("/users", func(r chi.Router) {
		r.Get("/", getAll(s.d.u))
		r.Get("/{id}", getById(s.d.u))
		r.Get("/{id}/address", s.getUserAddressById())
		r.Get("/{id}/company", s.getUserCompanyById())
		r.Get("/{id}/address/geo", s.getUserGeoById())
	})
}
