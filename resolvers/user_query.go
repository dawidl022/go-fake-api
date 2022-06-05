package resolvers

import (
	"strconv"

	"github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"

	"github.com/dawidl022/go-fake-api/models"
)

type UserQuery struct {
	db *gorm.DB
}

func NewUserQuery(db *gorm.DB) *UserQuery {
	return &UserQuery{db: db}
}

func (u *UserQuery) Users() []*User {
	var users []*models.User
	u.db.Preload("Address").Preload("Company").Find(&users)

	var resolvers []*User
	for _, user := range users {
		resolvers = append(resolvers, &User{u: user})
	}
	return resolvers
}

type userByIdArgs struct {
	ID graphql.ID
}

func (u *UserQuery) UserById(args userByIdArgs) (*User, error) {
	var user models.User
	id, err := strconv.Atoi(string(args.ID))
	if err != nil {
		return nil, err
	}
	err = u.db.Joins("Address").Joins("Company").First(&user, id).Error
	return &User{u: &user}, err
}
