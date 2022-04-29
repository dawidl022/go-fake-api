package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type RequestError int

func (e RequestError) Error() string {
	return "Request not completed"
}

type Album struct {
	Id     int
	UserId int
	Title  string
}

type Comment struct {
	Id     int
	PostId int
	Name   string
	Email  string
	Body   string
}

type Photo struct {
	Id           int
	AlbumId      int
	Title        string
	Url          string
	ThumbnailUrl string
}

type Post struct {
	Id     int
	UserId int
	Title  string
	Body   string
}

type Todo struct {
	Id        int
	UserId    int
	Title     string
	Completed bool
}

type User struct {
	Id       int
	Name     string
	Username string
	Email    string
	Address  *AddressDetails
	Phone    string
	Website  string
	Company  *CompanyDetails
}

type AddressDetails struct {
	Street  string
	Suite   string
	City    string
	Zipcode string
	Geo     *Coords
}

type Coords struct {
	Lat string
	Lng string
}

type CompanyDetails struct {
	Name        string
	CatchPhrase string
	Bs          string
}

type Data struct {
	a      []Album
	c      []Comment
	photos []Photo
	posts  []Post
	t      []Todo
	u      []User
}

type Server struct {
	r chi.Router
	d *Data
}

func loadData() *Data {
	d := Data{}

	rawAlbums, _ := os.ReadFile("data/albums.json")
	json.Unmarshal(rawAlbums, &d.a)

	rawComments, _ := os.ReadFile("data/comments.json")
	json.Unmarshal(rawComments, &d.c)

	rawPhotos, _ := os.ReadFile("data/photos.json")
	json.Unmarshal(rawPhotos, &d.photos)

	rawPosts, _ := os.ReadFile("data/posts.json")
	json.Unmarshal(rawPosts, &d.posts)

	rawTodos, _ := os.ReadFile("data/todos.json")
	json.Unmarshal(rawTodos, &d.t)

	rawUsers, _ := os.ReadFile("data/users.json")
	json.Unmarshal(rawUsers, &d.u)

	return &d
}

func NewServer() Server {
	s := Server{
		r: chi.NewRouter(),
		d: loadData(),
	}

	s.r.Use(middleware.RequestID)
	s.r.Use(middleware.RealIP)
	s.r.Use(middleware.Logger)
	s.r.Use(middleware.Recoverer)

	s.r.Use(middleware.Timeout(60 * time.Second))

	return s
}

func main() {
	s := NewServer()
	s.r.Use(JSONHeaders)

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

	http.ListenAndServe(":3000", s.r)
}

func JSONHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}

func getAll[T any](d []T) http.HandlerFunc {
	all, _ := json.Marshal(d)
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write(all)
	}
}

func getDataById[T any](d []T, w *http.ResponseWriter, r *http.Request) (T, error) {
	idParam := chi.URLParam(r, "id")
	i, err := checkIdParam(idParam, d, w)
	var zero T

	if err != nil {
		return zero, err
	}

	return d[i-1], nil
}

func getById[T any](d []T) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dat, err := getDataById(d, &w, r)

		if err != nil {
			return
		}

		res, _ := json.Marshal(dat)
		w.Write(res)
	}
}

// How to DRY up the following methods?

func (s *Server) getUserAddressById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dat, err := getDataById(s.d.u, &w, r)

		if err != nil {
			return
		}

		res, _ := json.Marshal(dat.Address)
		w.Write(res)
	}
}

func (s *Server) getUserCompanyById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dat, err := getDataById(s.d.u, &w, r)

		if err != nil {
			return
		}

		res, _ := json.Marshal(dat.Company)
		w.Write(res)
	}
}

func (s *Server) getUserGeoById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dat, err := getDataById(s.d.u, &w, r)

		if err != nil {
			return
		}

		res, _ := json.Marshal(dat.Address.Geo)
		w.Write(res)
	}
}

func (s *Server) getCommentsByPostId() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		i, err := checkIdParam(idParam, s.d.posts, &w)

		if err != nil {
			return
		}

		var dat []Comment

		for _, c := range s.d.c {
			if c.PostId == i {
				dat = append(dat, c)
			}
		}

		res, _ := json.Marshal(dat)
		w.Write(res)
	}
}

func checkIdParam[T any](id string, d []T, w *http.ResponseWriter) (int, error) {
	i, err := strconv.Atoi(id)

	if err != nil {
		http.Error(*w, http.StatusText(400), 400)
		return 0, RequestError(400)
	}

	if i > len(d) {
		http.Error(*w, http.StatusText(404), 404)
		return 0, RequestError(404)
	}

	return i, nil
}
