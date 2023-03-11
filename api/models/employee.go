package models

import "fmt"

type Employee struct {
	EmployeeID int    `json:"employee_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Title      string `json:"title"`
	WorkPhone  string `json:"work_phone"`
}

// example method function
func (c Employee) PrintEmployeeName() {
	fmt.Printf("This employee name is %v %v \n", c.FirstName, c.LastName)
}
