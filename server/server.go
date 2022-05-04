package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"server/config"
	"server/models"
	"server/resolvers"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	graphql "github.com/graph-gophers/graphql-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

	_, err = initDB(conf)
	if err != nil {
		log.Fatal("Cannot initialise database:", err)
	}

	root, err := resolvers.NewRootResolver(conf.BaseDir)
	if err != nil {
		log.Fatal("Cannot load data files:", err)
	}

	schema := graphql.MustParseSchema(string(b), root)
	s.routes(schema, conf.BaseDir)
}

func initDB(conf *config.Config) (*gorm.DB, error) {
	db, err := connectDB(conf)
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.Album{},
		&models.Post{},
	)
	if err != nil {
		return nil, err
	}

	err = seedDB(db, conf)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func seedDB(db *gorm.DB, conf *config.Config) error {
	var count int64
	db.Model(&models.Album{}).Count(&count)
	if count == 0 {
		err := load[models.Album](db, "albums", conf)
		if err != nil {
			return err
		}
	}

	db.Model(&models.Post{}).Count(&count)
	if count == 0 {
		err := load[models.Post](db, "posts", conf)
		if err != nil {
			return err
		}
	}

	return nil
}

func load[T any](db *gorm.DB, filename string, conf *config.Config) error {
	raw, err := os.ReadFile(fmt.Sprintf("%sdata/%s.json", conf.BaseDir, filename))
	if err != nil {
		return err
	}

	var models []*T
	err = json.Unmarshal(raw, &models)
	if err != nil {
		return err
	}

	db.Create(models)
	return nil
}

func connectDB(conf *config.Config) (*gorm.DB, error) {
	return gorm.Open(postgres.New(postgres.Config{
		DSN: conf.DatabaseUrl,
	}), &gorm.Config{})
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
