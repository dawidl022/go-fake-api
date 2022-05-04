package resolvers

import "gorm.io/gorm"

type RootResolver struct {
	*AlbumQuery
	*PostQuery
}

func NewRootResolver(db *gorm.DB) *RootResolver {

	return &RootResolver{
		AlbumQuery: NewAlbumQuery(db),
		PostQuery:  NewPostQuery(db),
	}
}
