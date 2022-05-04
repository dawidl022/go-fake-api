package resolvers

import (
	"server/models"

	"github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

type PostQuery struct {
	db *gorm.DB
}

func NewPostQuery(db *gorm.DB) *PostQuery {
	return &PostQuery{db: db}
}

func (p *PostQuery) Posts() []*Post {
	var posts []*models.Post
	p.db.Find(&posts)
	return makePostResolvers(posts)
}

type postArgs struct {
	ID graphql.ID
}

func (p *PostQuery) Post(args postArgs) *Post {
	var post models.Post
	p.db.First(&post, args.ID)
	return &Post{pm: &post}
}

type postsByUserArgs struct {
	UserID graphql.ID
}

func (p *PostQuery) PostsByUser(args postsByUserArgs) []*Post {
	var posts []*models.Post
	p.db.Where("user_id = ?", args.UserID).Find(&posts)
	return makePostResolvers(posts)
}

func makePostResolvers(posts []*models.Post) []*Post {
	var resolvers []*Post
	for _, post := range posts {
		resolvers = append(resolvers, &Post{pm: post})
	}
	return resolvers
}
