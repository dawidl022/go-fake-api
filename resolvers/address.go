package resolvers

import "github.com/dawidl022/go-fake-api/models"

type Address struct {
	a *models.Address
}

type Geo struct {
	g *models.Geo
}

func (a *Address) Street() string {
	return a.a.Street
}

func (a *Address) Suite() *string {
	return a.a.Suite
}

func (a *Address) City() string {
	return a.a.City
}

func (a *Address) Zipcode() string {
	return a.a.Zipcode
}

func (a *Address) Country() string {
	return a.a.Country
}

func (a *Address) Geo() *Geo {
	return &Geo{g: a.a.Geo}
}

func (g *Geo) Lat() string {
	return g.g.Lat
}

func (g *Geo) Lng() string {
	return g.g.Lng
}
