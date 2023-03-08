package models

import "fmt"

type Product struct {
	ProductID   int
	ProductName string
	UnitPrice   int
	InStock     string
}

func (c Product) PrintProductName() {
	fmt.Printf("This product name is %v \n", c.ProductName)
}
