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
	OrderID             int                    `json:"order_id"`
	OrderDate           string                 `json:"order_date"`
	PurchaseOrderNumber string                 `json:"purchase_order_number"`
	ShipDate            string                 `json:"ship_date"`
	FreightCharge       int                    `json:"freight_charge"`
	Taxes               int                    `json:"taxes"`
	PaymentReceived     string                 `json:"payment_received"`
	Comment             string                 `json:"comment"`
	TotalPrice          int                    `json:"total_price"`
	FinalTotalPrice     int                    `json:"final_total_price"`
	ShippingMethod      ShippingMethod         `json:"shipping_method"`
	Employee            Employee               `json:"employee"`
	Customer            Customer               `json:"customer"`
	Details             []OrderDetailsResponse `json:"details"`
}

type OrderDetailsResponse struct {
	OrderDetailID int     `json:"order_detail_id"`
	Quantity      int     `json:"quantity"`
	UnitPrice     int     `json:"unit_price"`
	Discount      int     `json:"discount"`
	Price         int     `json:"price"`
	Product       Product `json:"product"`
}

// example method function
func (c Order) PrintOrderDate() {
	fmt.Printf("This order has ordered in date %v \n", c.OrderDate.String())
}

// example method function
func (c OrderDetails) PrintOrderDetailsPrice() {
	fmt.Printf("This order price is is %v \n", c.UnitPrice)
}
