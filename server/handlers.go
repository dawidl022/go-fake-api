package server

// import (
// 	"encoding/json"
// 	"net/http"
// 	"server/models"
// 	"strconv"

// 	"github.com/go-chi/chi/v5"
// )

// type requestError int

// func (e requestError) Error() string {
// 	return "Request not completed"
// }

// func getAll[T any](d []T) http.HandlerFunc {
// 	all, _ := json.Marshal(d)
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		w.Write(all)
// 	}
// }

// func checkIdParam[T any](id string, d []T, w *http.ResponseWriter) (int, error) {
// 	i, err := strconv.Atoi(id)

// 	if err != nil {
// 		http.Error(*w, http.StatusText(400), 400)
// 		return 0, requestError(400)
// 	}

// 	if i > len(d) {
// 		http.Error(*w, http.StatusText(404), 404)
// 		return 0, requestError(404)
// 	}

// 	return i, nil
// }

// func getDataById[T any](d []T, w *http.ResponseWriter, r *http.Request) (T, error) {
// 	idParam := chi.URLParam(r, "id")
// 	i, err := checkIdParam(idParam, d, w)
// 	var zero T

// 	if err != nil {
// 		return zero, err
// 	}

// 	return d[i-1], nil
// }

// func getById[T any](d []T) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		dat, err := getDataById(d, &w, r)

// 		if err != nil {
// 			return
// 		}

// 		res, _ := json.Marshal(dat)
// 		w.Write(res)
// 	}
// }

// // How to DRY up the following methods?

// func (s *server) getUserAddressById() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		dat, err := getDataById(s.d.u, &w, r)

// 		if err != nil {
// 			return
// 		}

// 		res, _ := json.Marshal(dat.Address)
// 		w.Write(res)
// 	}
// }

// func (s *server) getUserCompanyById() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		dat, err := getDataById(s.d.u, &w, r)

// 		if err != nil {
// 			return
// 		}

// 		res, _ := json.Marshal(dat.Company)
// 		w.Write(res)
// 	}
// }

// func (s *server) getUserGeoById() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		dat, err := getDataById(s.d.u, &w, r)

// 		if err != nil {
// 			return
// 		}

// 		res, _ := json.Marshal(dat.Address.Geo)
// 		w.Write(res)
// 	}
// }

// func (s *server) getCommentsByPostId() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		idParam := chi.URLParam(r, "id")
// 		i, err := checkIdParam(idParam, s.d.posts, &w)

// 		if err != nil {
// 			return
// 		}

// 		var dat []models.Comment

// 		for _, c := range s.d.c {
// 			if c.PostId == i {
// 				dat = append(dat, c)
// 			}
// 		}

// 		res, _ := json.Marshal(dat)
// 		w.Write(res)
// 	}
// }
