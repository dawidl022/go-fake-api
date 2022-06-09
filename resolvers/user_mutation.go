package resolvers

import (
	"github.com/dawidl022/go-fake-api/models"
	"gorm.io/gorm"
)

type UserMutation struct {
	db *gorm.DB
}

func NewUserMutation(db *gorm.DB) *UserMutation {
	return &UserMutation{db: db}
}

type createUserArgs struct {
	User *struct {
		Name     string
		Username string
		Email    string
		Address  *struct {
			Street  string
			Suite   *string
			City    string
			Zipcode string
			Country string
		}
		Phone   *string
		Website *string
		Company *struct {
			Name        string
			CatchPhrase *string
			Bs          *string
		}
	}
}

func (u *UserMutation) CreateUser(args createUserArgs) (*User, error) {
	var address *models.Address
	if args.User.Address != nil {
		address = &models.Address{
			Street:  args.User.Address.Street,
			Suite:   args.User.Address.Suite,
			City:    args.User.Address.City,
			Zipcode: args.User.Address.Zipcode,
			Country: args.User.Address.Country,
		}
	}
	var company *models.Company
	if args.User.Company != nil {
		company = &models.Company{
			Name:        args.User.Company.Name,
			CatchPhrase: args.User.Company.CatchPhrase,
			Bs:          args.User.Company.Bs,
		}
	}

	user := models.User{
		Name:     args.User.Name,
		Username: args.User.Username,
		Email:    args.User.Email,
		Address:  address,
		Phone:    args.User.Phone,
		Website:  args.User.Website,
		Company:  company,
	}
	return &User{u: &user}, u.db.Create(&user).Error
}
