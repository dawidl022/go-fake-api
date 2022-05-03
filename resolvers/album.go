package resolvers

import (
	"fmt"
	"server/models"

	"github.com/graph-gophers/graphql-go"
)

type Album struct {
	am *models.Album
}

func (a *Album) Id() graphql.ID {
	return graphql.ID(fmt.Sprint(a.am.Id))
}

func (a *Album) UserId() graphql.ID {
	return graphql.ID(fmt.Sprint(a.am.UserId))
}

func (a *Album) Title() string {
	return a.am.Title
}
