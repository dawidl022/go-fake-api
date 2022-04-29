package models

type Comment struct {
	Id     int
	PostId int
	Name   string
	Email  string
	Body   string
}
