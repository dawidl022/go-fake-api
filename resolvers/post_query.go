package resolvers

import (
	"encoding/json"
	"os"
	"server/models"
	"strconv"

	"github.com/graph-gophers/graphql-go"
)

type PostQuery struct {
	posts map[string]*Post
	keys  []string
}

func NewPostQuery() *PostQuery {
	p := PostQuery{}
	p.setup()

	return &p
}

func (p *PostQuery) setup() error {
	rawPosts, err := os.ReadFile("data/posts.json")
	if err != nil {
		return err
	}

	var pm []*models.Post
	err = json.Unmarshal(rawPosts, &pm)
	if err != nil {
		return err
	}

	p.posts = make(map[string]*Post)
	for _, post := range pm {
		id := strconv.Itoa(post.Id)
		p.posts[id] = &Post{pm: post}
		p.keys = append(p.keys, id)
	}

	return nil
}

func (p *PostQuery) Posts() []*Post {
	res := make([]*Post, 0, len(p.posts))

	for _, k := range p.keys {
		res = append(res, p.posts[k])
	}

	return res
}

type postArgs struct {
	ID graphql.ID
}

func (p *PostQuery) Post(args postArgs) *Post {
	// TODO handle invalid id

	return p.posts[string(args.ID)]
}

type postsByUserArgs struct {
	UserID graphql.ID
}

func (p *PostQuery) PostsByUser(args postsByUserArgs) ([]*Post, error) {
	var res []*Post

	for _, k := range p.keys {
		userId, err := strconv.Atoi(string(args.UserID))
		if err != nil {
			return nil, err
		}

		p := p.posts[k]
		if p.pm.UserId == userId {
			res = append(res, p)
		}
	}

	return res, nil
}
