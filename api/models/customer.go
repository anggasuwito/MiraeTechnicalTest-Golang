package models

import "fmt"

type Customer struct {
	CustomerID      int    `json:"customer_id"`
	CompanyName     string `json:"company_name"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	BillingAddress  string `json:"billing_address"`
	City            string `json:"city"`
	Province        string `json:"province"`
	ZipCode         string `json:"zip_code"`
	Email           string `json:"email"`
	CompanyWebsite  string `json:"company_website"`
	PhoneNumber     string `json:"phone_number"`
	FaxNumber       string `json:"fax_number"`
	ShipAddress     string `json:"ship_address"`
	ShipCity        string `json:"ship_city"`
	ShipProvince    string `json:"ship_province"`
	ShipZipCode     string `json:"ship_zip_code"`
	ShipPhoneNumber string `json:"ship_phone_number"`
}

// example method function
func (c Customer) PrintCustomerName() {
	fmt.Printf("This customer name is %v %v \n", c.FirstName, c.LastName)
}
