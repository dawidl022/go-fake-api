package resolvers

import (
	"strconv"

	"github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"

	"github.com/dawidl022/go-fake-api/models"
)

type AlbumMutation struct {
	db    *gorm.DB
	query *AlbumQuery
}

func NewAlbumMutation(db *gorm.DB) *AlbumMutation {
	return &AlbumMutation{db: db, query: &AlbumQuery{db: db}}
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

	album := models.Album{UserId: int32(userId), Title: args.Title}
	a.db.Create(&album)

	return &Album{am: &album}, nil
}

type deleteAlbumArgs struct {
	ID graphql.ID
}

func (a *AlbumMutation) DeleteAlbum(args deleteAlbumArgs) (*Album, error) {
	album, err := a.query.Album(albumArgs(args))
	if err != nil {
		return nil, err
	}

	a.db.Delete(&models.Album{}, args.ID)
	return album, nil
}

type updateAlbumArgs struct {
	ID     graphql.ID
	UserId *int32
	Title  *string
}

func (a *AlbumMutation) UpdateAlbum(args updateAlbumArgs) (*Album, error) {
	album, err := a.query.Album(albumArgs{ID: args.ID})
	if err != nil {
		return nil, err
	}

	if args.UserId != nil {
		album.am.UserId = *args.UserId
	}
	if args.Title != nil {
		album.am.Title = *args.Title
	}

	a.db.Save(&album.am)
	return album, nil
}
