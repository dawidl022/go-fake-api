package models

type User struct {
	Id       int
	Name     string
	Username string
	Email    string
	Address  *AddressDetails
	Phone    string
	Website  string
	Company  *CompanyDetails
}

type AddressDetails struct {
	Street  string
	Suite   string
	City    string
	Zipcode string
	Geo     *Coords
}

type Coords struct {
	Lat string
	Lng string
}

type CompanyDetails struct {
	Name        string
	CatchPhrase string
	Bs          string
}
