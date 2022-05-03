package resolvers

type RootResolver struct {
	*AlbumQuery
	*PostQuery
}

func NewRootResolver() *RootResolver {
	return &RootResolver{
		AlbumQuery: NewAlbumQuery(),
		PostQuery:  NewPostQuery(),
	}
}
