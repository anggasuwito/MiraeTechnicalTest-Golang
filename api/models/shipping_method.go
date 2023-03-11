package models

import "fmt"

type ShippingMethod struct {
	ShippingMethodID int    `json:"shipping_method_id"`
	Method           string `json:"method"`
}

// example method function
func (c ShippingMethod) PrintShippingMethod() {
	fmt.Printf("This shipping method is %v \n", c.Method)
}
