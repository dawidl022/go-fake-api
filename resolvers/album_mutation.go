package resolvers

import (
	"server/models"
	"strconv"

	"github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

type AlbumMutation struct {
	db *gorm.DB
}

func NewAlbumMutation(db *gorm.DB) *AlbumMutation {
	return &AlbumMutation{db: db}
}

type createAlbumArgs struct {
	UserId graphql.ID
	Title  string
}

func (a *AlbumMutation) CreateAlbum(args createAlbumArgs) (*Album, error) {
	// TODO add user resolvers to check if user is valid
	// var user models.User
	// result := a.db.First(&user, args.UserId)

	// if result.Error != nil {
	// 	return nil, result.Error
	// }

	userId, err := strconv.Atoi(string(args.UserId))
	if err != nil {
		return nil, err
	}

	album := models.Album{UserId: userId, Title: args.Title}
	a.db.Create(&album)

	return &Album{am: &album}, nil
}
