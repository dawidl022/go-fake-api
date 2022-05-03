package resolvers

import (
	"encoding/json"
	"os"
	"server/models"
	"strconv"

	"github.com/graph-gophers/graphql-go"
)

type AlbumQuery struct {
	albums map[string]*Album
	keys   []string
}

func NewAlbumQuery() *AlbumQuery {
	a := AlbumQuery{}
	a.setup()

	return &a
}

func (a *AlbumQuery) setup() error {
	rawAlbums, err := os.ReadFile("data/albums.json")
	if err != nil {
		return err
	}

	var am []*models.Album
	err = json.Unmarshal(rawAlbums, &am)
	if err != nil {
		return err
	}

	a.albums = make(map[string]*Album)
	for _, album := range am {
		id := strconv.Itoa(album.Id)
		a.albums[id] = &Album{am: album}
		a.keys = append(a.keys, id)
	}

	return nil
}

func (a *AlbumQuery) Albums() []*Album {
	res := make([]*Album, 0, len(a.albums))

	for _, k := range a.keys {
		res = append(res, a.albums[k])
	}

	return res
}

type albumArgs struct {
	ID graphql.ID
}

func (a *AlbumQuery) Album(args albumArgs) *Album {
	// TODO handle invalid id

	return a.albums[string(args.ID)]
}

type albumsByUserArgs struct {
	UserID graphql.ID
}

func (a *AlbumQuery) AlbumsByUser(args albumsByUserArgs) ([]*Album, error) {
	var res []*Album

	for _, k := range a.keys {
		userId, err := strconv.Atoi(string(args.UserID))
		if err != nil {
			return nil, err
		}

		a := a.albums[k]
		if a.am.UserId == userId {
			res = append(res, a)
		}
	}

	return res, nil
}
