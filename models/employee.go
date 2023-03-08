package models

import "fmt"

type Employee struct {
	EmployeeID int
	FirstName  string
	LastName   string
	Title      string
	WorkPhone  string
}

func (c Employee) PrintEmployeeName() {
	fmt.Printf("This employee name is %v %v \n", c.FirstName, c.LastName)
}
