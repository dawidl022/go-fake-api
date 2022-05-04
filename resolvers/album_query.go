package resolvers

import (
	"server/models"

	"github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

type AlbumQuery struct {
	db *gorm.DB
}

func NewAlbumQuery(db *gorm.DB) *AlbumQuery {
	return &AlbumQuery{db: db}
}

func (a *AlbumQuery) Albums() []*Album {
	var albums []*models.Album
	a.db.Find(&albums)
	return makeAlbumResolvers(albums)
}

type albumArgs struct {
	ID graphql.ID
}

func (a *AlbumQuery) Album(args albumArgs) (*Album, error) {
	var album models.Album
	result := a.db.First(&album, args.ID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &Album{am: &album}, nil
}

type albumsByUserArgs struct {
	UserID graphql.ID
}

func (a *AlbumQuery) AlbumsByUser(args albumsByUserArgs) []*Album {
	var albums []*models.Album
	a.db.Where("user_id = ?", args.UserID).Find(&albums)
	return makeAlbumResolvers(albums)
}

func makeAlbumResolvers(albums []*models.Album) []*Album {
	var resolvers []*Album
	for _, album := range albums {
		resolvers = append(resolvers, &Album{am: album})
	}
	return resolvers
}
