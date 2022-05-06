package resolvers

import (
	"server/models"
	"strconv"

	"github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

type PostMutation struct {
	db    *gorm.DB
	query *PostQuery
}

func NewPostMutation(db *gorm.DB) *PostMutation {
	return &PostMutation{db: db, query: &PostQuery{db: db}}
}

type createPostArgs struct {
	UserId graphql.ID
	Title  string
	Body   string
}

func (p *PostMutation) CreatePost(args createPostArgs) (*Post, error) {
	userId, err := strconv.Atoi(string(args.UserId))
	if err != nil {
		return nil, err
	}
	post := models.Post{UserId: int32(userId), Title: args.Title, Body: args.Body}

	p.db.Create(&post)
	return &Post{pm: &post}, nil
}

type deletePostArgs struct {
	ID graphql.ID
}

func (p *PostMutation) DeletePost(args deletePostArgs) (*Post, error) {
	post, err := p.query.Post(postArgs(args))
	if err != nil {
		return nil, err
	}

	p.db.Delete(&models.Post{}, args.ID)
	return post, nil
}

type updatePostArgs struct {
	ID     graphql.ID
	UserId *int32
	Title  *string
	Body   *string
}

func (p *PostMutation) UpdatePost(args updatePostArgs) (*Post, error) {
	post, err := p.query.Post(postArgs{ID: args.ID})
	if err != nil {
		return nil, err
	}

	if args.UserId != nil {
		post.pm.UserId = *args.UserId
	}
	if args.Title != nil {
		post.pm.Title = *args.Title
	}
	if args.Body != nil {
		post.pm.Body = *args.Body
	}

	p.db.Save(&post.pm)
	return post, nil
}
