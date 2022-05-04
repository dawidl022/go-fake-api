package resolvers

import (
	"server/models"
	"strconv"

	"github.com/graph-gophers/graphql-go"
)

type Album struct {
	am *models.Album
}

func (a *Album) ID() graphql.ID {
	return graphql.ID(strconv.Itoa(a.am.ID))
}

func (a *Album) UserId() graphql.ID {
	return graphql.ID(strconv.Itoa(a.am.UserId))
}

func (a *Album) Title() string {
	return a.am.Title
}
