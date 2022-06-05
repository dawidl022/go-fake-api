package resolvers

import (
	"strconv"

	"github.com/graph-gophers/graphql-go"

	"github.com/dawidl022/go-fake-api/models"
)

type Album struct {
	am *models.Album
}

func (a *Album) ID() graphql.ID {
	return graphql.ID(strconv.Itoa(int(a.am.ID)))
}

func (a *Album) UserId() graphql.ID {
	return graphql.ID(strconv.Itoa(int(a.am.UserId)))
}

func (a *Album) Title() string {
	return a.am.Title
}
