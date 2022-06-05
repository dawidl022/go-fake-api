package resolvers

import (
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
	u.db.Find(&users)

	var resolvers []*User
	for _, user := range users {
		resolvers = append(resolvers, &User{u: user})
	}
	return resolvers
}
