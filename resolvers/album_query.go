package resolvers

import (
	"encoding/json"
	"fmt"
	"os"
	"server/models"
	"strconv"

	"github.com/graph-gophers/graphql-go"
)

type AlbumQuery struct {
	a    map[string]*Album
	keys []string
}

func (aq *AlbumQuery) Setup() {
	rawAlbums, _ := os.ReadFile("data/albums.json")
	var am []*models.Album

	json.Unmarshal(rawAlbums, &am)
	aq.a = make(map[string]*Album)

	for _, album := range am {
		id := fmt.Sprint(album.Id)
		aq.a[id] = &Album{am: album}
		aq.keys = append(aq.keys, id)
	}
}

func (aq *AlbumQuery) Albums() []*Album {
	res := make([]*Album, 0, len(aq.a))

	for _, k := range aq.keys {
		res = append(res, aq.a[k])
	}

	return res
}

func (aq *AlbumQuery) Album(args struct{ Id graphql.ID }) *Album {
	// TODO handle out of bounds error

	return aq.a[string(args.Id)]
}

func (aq *AlbumQuery) AlbumsByUser(args struct{ UserId graphql.ID }) []*Album {
	var res []*Album

	for _, k := range aq.keys {
		userId, _ := strconv.Atoi(string(args.UserId))
		a := aq.a[k]
		if a.am.UserId == userId {
			res = append(res, a)
		}
	}

	return res
}
