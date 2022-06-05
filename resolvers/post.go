package resolvers

import (
	"strconv"

	"github.com/graph-gophers/graphql-go"

	"github.com/dawidl022/go-fake-api/models"
)

type Post struct {
	pm *models.Post
}

func (p *Post) Id() graphql.ID {
	return graphql.ID(strconv.Itoa(int(p.pm.Id)))
}

func (p *Post) UserId() graphql.ID {
	return graphql.ID(strconv.Itoa(int(p.pm.UserId)))
}

func (p *Post) Title() string {
	return p.pm.Title
}

func (p *Post) Body() string {
	return p.pm.Body
}
