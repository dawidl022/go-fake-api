package resolvers

import (
	"encoding/json"
	"os"
	"server/models"
)

type Query struct{}

func (*Query) Albums() []*Album {
	rawAlbums, _ := os.ReadFile("data/albums.json")
	var am []*models.Album

	json.Unmarshal(rawAlbums, &am)

	var a []*Album
	for _, album := range am {
		a = append(a, &Album{am: album})
	}

	return a
}
