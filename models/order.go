package models

import (
	"fmt"
	"time"
)

type Order struct {
	OrderID             int
	CustomerID          int
	EmployeeID          int
	OrderDate           time.Time
	PurchaseOrderNumber string
	ShipDate            time.Time
	ShippingMethodID    int
	FreightCharge       int
	Taxes               int
	PaymentReceived     string
	Comment             string
}

func (c Order) PrintOrderDate() {
	fmt.Printf("This order has ordered in date %v \n", c.OrderDate.String())
}
