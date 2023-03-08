package models

import "fmt"

type ShippingMethod struct {
	ShippingMethodID int
	Method           string
}

func (c ShippingMethod) PrintShippingMethod() {
	fmt.Printf("This shipping method is %v \n", c.Method)
}
