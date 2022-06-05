package resolvers

import "gorm.io/gorm"

type RootResolver struct {
	*AlbumQuery
	*AlbumMutation
	*PostQuery
	*PostMutation
	*UserQuery
}

func NewRootResolver(db *gorm.DB) *RootResolver {

	return &RootResolver{
		AlbumQuery:    NewAlbumQuery(db),
		AlbumMutation: NewAlbumMutation(db),
		PostQuery:     NewPostQuery(db),
		PostMutation:  NewPostMutation(db),
		UserQuery:     NewUserQuery(db),
	}
}
