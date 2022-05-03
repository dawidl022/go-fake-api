package resolvers

import (
	"encoding/json"
	"os"
	"server/models"
	"strconv"

	"github.com/graph-gophers/graphql-go"
)

type AlbumQuery struct {
	a []*Album
	// TODO change to map[string]*Album
}

func (aq *AlbumQuery) Setup() {
	rawAlbums, _ := os.ReadFile("data/albums.json")
	var am []*models.Album

	json.Unmarshal(rawAlbums, &am)

	for _, album := range am {
		aq.a = append(aq.a, &Album{am: album})
	}
}

func (aq *AlbumQuery) Albums() []*Album {
	return aq.a
}

func (aq *AlbumQuery) Album(args struct{ Id graphql.ID }) *Album {
	// TODO handle out of bounds error

	i, _ := strconv.Atoi(string(args.Id))
	return aq.a[i-1]
}

func (aq *AlbumQuery) AlbumByUser(args struct{ UserId graphql.ID }) []*Album {
	var res []*Album

	for _, a := range aq.a {
		userId, _ := strconv.Atoi(string(args.UserId))
		if a.am.UserId == userId {
			res = append(res, a)
		}
	}

	return res
}
