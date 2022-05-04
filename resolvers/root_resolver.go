package resolvers

type RootResolver struct {
	*AlbumQuery
	*PostQuery
}

func NewRootResolver(basedir string) (*RootResolver, error) {
	a, err := NewAlbumQuery(basedir)
	if err != nil {
		return nil, err
	}

	p, err := NewPostQuery(basedir)
	if err != nil {
		return nil, err
	}

	return &RootResolver{
		AlbumQuery: a,
		PostQuery:  p,
	}, nil
}
