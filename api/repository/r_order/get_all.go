package r_order

import (
	"MiraeTechnicalTest-Golang/api/models"
)

func (r OrderRepo) GetAllOrderRepo(offset int, limit int) (res []models.OrderResponse, err error) {
	rows, err := r.DB.
		Query(`
			SELECT 
			    od.order_detail_id,od.quantity,od.unit_price,od.discount,
			    p.product_id,p.product_name,p.unit_price,p.in_stock,
			    o.order_id,o.order_date,o.purchase_order_number,o.ship_date,o.freight_charge,o.taxes,o.payment_received,o.comment,
			    sm.shipping_method_id,sm.shipping_method,
			    e.employee_id,e.first_name,e.last_name,e.title,e.work_phone,
			    c.customer_id,c.company_name,c.first_name,c.last_name,c.billing_address,c.city,c.province,c.zip_code,c.email,c.company_website,c.phone_number,c.fax_number,c.ship_address,c.ship_city,c.ship_province,c.ship_zip_code,c.ship_phone_number
			FROM order_details od
			LEFT JOIN products p ON p.product_id = od.product_id
			LEFT JOIN orders o ON o.order_id = od.order_id
			LEFT JOIN shipping_methods sm ON sm.shipping_method_id = o.shipping_method_id
			LEFT JOIN employees e ON e.employee_id = o.employee_id
			LEFT JOIN customers c ON c.customer_id = o.customer_id
			LIMIT ? OFFSET ? 
		`, limit, offset)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var ord models.OrderResponse
		err = rows.Scan(
			&ord.OrderDetailID, &ord.Quantity, &ord.UnitPrice, &ord.Discount,
			&ord.Product.ProductID, &ord.Product.ProductName, &ord.Product.UnitPrice, &ord.Product.InStock,
			&ord.Order.OrderID, &ord.Order.OrderDate, &ord.Order.PurchaseOrderNumber, &ord.Order.ShipDate, &ord.Order.FreightCharge, &ord.Order.Taxes, &ord.Order.PaymentReceived, &ord.Order.Comment,
			&ord.Order.ShippingMethod.ShippingMethodID, &ord.Order.ShippingMethod.Method,
			&ord.Order.Employee.EmployeeID, &ord.Order.Employee.FirstName, &ord.Order.Employee.LastName, &ord.Order.Employee.Title, &ord.Order.Employee.WorkPhone,
			&ord.Order.Customer.CustomerID, &ord.Order.Customer.CompanyName, &ord.Order.Customer.FirstName, &ord.Order.Customer.LastName, &ord.Order.Customer.BillingAddress, &ord.Order.Customer.City, &ord.Order.Customer.Province, &ord.Order.Customer.ZipCode, &ord.Order.Customer.Email, &ord.Order.Customer.CompanyWebsite, &ord.Order.Customer.PhoneNumber, &ord.Order.Customer.FaxNumber, &ord.Order.Customer.ShipAddress, &ord.Order.Customer.ShipCity, &ord.Order.Customer.ShipProvince, &ord.Order.Customer.ShipZipCode, &ord.Order.Customer.ShipPhoneNumber,
		)
		if err != nil {
			return res, err
		}
		res = append(res, ord)
	}
	return res, err
}
