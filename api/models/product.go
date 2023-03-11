package models

import "fmt"

type Product struct {
	ProductID   int    `json:"product_id"`
	ProductName string `json:"product_name"`
	UnitPrice   int    `json:"unit_price"`
	InStock     string `json:"in_stock"`
}

// example method function
func (c Product) PrintProductName() {
	fmt.Printf("This product name is %v \n", c.ProductName)
}
