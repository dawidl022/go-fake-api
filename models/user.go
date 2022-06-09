package models

type User struct {
	Id        int
	Name      string
	Username  string
	Email     string
	AddressID *int
	Address   *Address
	Phone     *string
	Website   *string
	CompanyID *int
	Company   *Company
}

type Address struct {
	ID      int
	Street  string
	Suite   *string
	City    string
	Zipcode string
	Country string
	Geo     *Geo `gorm:"embedded"`
}

type Geo struct {
	Lat string
	Lng string
}

type Company struct {
	ID          int
	Name        string
	CatchPhrase *string
	Bs          *string
}
