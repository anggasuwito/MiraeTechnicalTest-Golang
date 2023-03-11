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

type OrderDetails struct {
	OrderDetailID int
	OrderID       int
	ProductID     int
	Quantity      int
	UnitPrice     int
	Discount      int
}

type OrderResponse struct {
	OrderDetailID int     `json:"order_detail_id"`
	Quantity      int     `json:"quantity"`
	UnitPrice     int     `json:"unit_price"`
	Discount      int     `json:"discount"`
	Product       Product `json:"product"`
	Order         struct {
		OrderID             int            `json:"order_id"`
		OrderDate           string         `json:"order_date"`
		PurchaseOrderNumber string         `json:"purchase_order_number"`
		ShipDate            string         `json:"ship_date"`
		FreightCharge       int            `json:"freight_charge"`
		Taxes               int            `json:"taxes"`
		PaymentReceived     string         `json:"payment_received"`
		Comment             string         `json:"comment"`
		ShippingMethod      ShippingMethod `json:"shipping_method"`
		Employee            Employee       `json:"employee"`
		Customer            Customer       `json:"customer"`
	} `json:"order"`
}

// example method function
func (c Order) PrintOrderDate() {
	fmt.Printf("This order has ordered in date %v \n", c.OrderDate.String())
}

// example method function
func (c OrderDetails) PrintOrderDetailsPrice() {
	fmt.Printf("This order price is is %v \n", c.UnitPrice)
}
