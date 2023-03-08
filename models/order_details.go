package models

import "fmt"

type OrderDetails struct {
	OrderDetailID int
	OrderID       int
	ProductID     int
	Quantity      int
	UnitPrice     int
	Discount      int
}

func (c OrderDetails) PrintOrderDetailsPrice() {
	fmt.Printf("This order price is is %v \n", c.UnitPrice)
}
