package resolvers

import "github.com/dawidl022/go-fake-api/models"

type Company struct {
	c *models.Company
}

func (c *Company) Name() string {
	return c.c.Name
}

func (c *Company) CatchPhrase() string {
	return c.c.CatchPhrase
}

func (c *Company) Bs() string {
	return c.c.Bs
}
