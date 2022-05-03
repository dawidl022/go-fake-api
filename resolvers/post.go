package resolvers

import (
	"fmt"
	"server/models"

	"github.com/graph-gophers/graphql-go"
)

type Post struct {
	pm *models.Post
}

func (p *Post) Id() graphql.ID {
	return graphql.ID(fmt.Sprint(p.pm.Id))
}

func (p *Post) UserId() graphql.ID {
	return graphql.ID(fmt.Sprint(p.pm.UserId))
}

func (p *Post) Title() string {
	return p.pm.Title
}

func (p *Post) Body() string {
	return p.pm.Body
}
