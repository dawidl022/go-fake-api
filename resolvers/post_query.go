package resolvers

import (
	"encoding/json"
	"fmt"
	"os"
	"server/models"
	"strconv"

	"github.com/graph-gophers/graphql-go"
)

type PostQuery struct {
	p    map[string]*Post
	keys []string
}

func (pq *PostQuery) setup() {
	rawPosts, _ := os.ReadFile("data/posts.json")
	var pm []*models.Post

	json.Unmarshal(rawPosts, &pm)
	pq.p = make(map[string]*Post)

	for _, post := range pm {
		id := fmt.Sprint(post.Id)
		pq.p[id] = &Post{pm: post}
		pq.keys = append(pq.keys, id)
	}
}

func (pq *PostQuery) Posts() []*Post {
	res := make([]*Post, 0, len(pq.p))

	for _, k := range pq.keys {
		res = append(res, pq.p[k])
	}

	return res
}

func (pq *PostQuery) Post(args struct{ Id graphql.ID }) *Post {
	// TODO handle out of bounds error

	return pq.p[string(args.Id)]
}

func (pq *PostQuery) PostsByUser(args struct{ UserId graphql.ID }) []*Post {
	var res []*Post

	for _, k := range pq.keys {
		userId, _ := strconv.Atoi(string(args.UserId))
		p := pq.p[k]
		if p.pm.UserId == userId {
			res = append(res, p)
		}
	}

	return res
}

func NewPostQuery() *PostQuery {
	pq := PostQuery{}
	pq.setup()

	return &pq
}
