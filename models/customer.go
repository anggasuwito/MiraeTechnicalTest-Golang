package models

import "fmt"

type Customer struct {
	CustomerID      int
	CompanyName     string
	FirstName       string
	LastName        string
	BillingAddress  string
	City            string
	Province        string
	ZipCode         string
	Email           string
	CompanyWebsite  string
	PhoneNumber     string
	FaxNumber       string
	ShipAddress     string
	ShipCity        string
	ShipProvince    string
	ShipZipCode     string
	ShipPhoneNumber string
}

func (c Customer) PrintCustomerName() {
	fmt.Printf("This customer name is %v %v \n", c.FirstName, c.LastName)
}
