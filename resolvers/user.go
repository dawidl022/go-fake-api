package resolvers

import (
	"strconv"

	"github.com/graph-gophers/graphql-go"

	"github.com/dawidl022/go-fake-api/models"
)

type User struct {
	u *models.User
}

func (u *User) Id() graphql.ID {
	return graphql.ID(strconv.Itoa(u.u.Id))
}

func (u *User) Name() string {
	return u.u.Name
}

func (u *User) Username() string {
	return u.u.Username
}

func (u *User) Email() string {
	return u.u.Email
}

func (u *User) Address() *Address {
	return &Address{a: u.u.Address}
}

func (u *User) Phone() *string {
	return u.u.Phone
}

func (u *User) Website() *string {
	return u.u.Phone
}

func (u *User) Company() *Company {
	return &Company{c: u.u.Company}
}
