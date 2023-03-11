package r_order

import (
	"MiraeTechnicalTest-Golang/api/models"
)

func (r OrderRepo) GetAllOrderRepo(offset int, limit int) (res []models.OrderResponse, err error) {
	rows, err := r.DB.
		Query(`
			SELECT 
			    o.order_id,o.order_date,o.purchase_order_number,o.ship_date,o.freight_charge,o.taxes,o.payment_received,o.comment,
			    sm.shipping_method_id,sm.shipping_method,
			    e.employee_id,e.first_name,e.last_name,e.title,e.work_phone,
			    c.customer_id,c.company_name,c.first_name,c.last_name,c.billing_address,c.city,c.province,c.zip_code,c.email,c.company_website,c.phone_number,c.fax_number,c.ship_address,c.ship_city,c.ship_province,c.ship_zip_code,c.ship_phone_number
			FROM orders o
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
			&ord.OrderID, &ord.OrderDate, &ord.PurchaseOrderNumber, &ord.ShipDate, &ord.FreightCharge, &ord.Taxes, &ord.PaymentReceived, &ord.Comment,
			&ord.ShippingMethod.ShippingMethodID, &ord.ShippingMethod.Method,
			&ord.Employee.EmployeeID, &ord.Employee.FirstName, &ord.Employee.LastName, &ord.Employee.Title, &ord.Employee.WorkPhone,
			&ord.Customer.CustomerID, &ord.Customer.CompanyName, &ord.Customer.FirstName, &ord.Customer.LastName, &ord.Customer.BillingAddress, &ord.Customer.City, &ord.Customer.Province, &ord.Customer.ZipCode, &ord.Customer.Email, &ord.Customer.CompanyWebsite, &ord.Customer.PhoneNumber, &ord.Customer.FaxNumber, &ord.Customer.ShipAddress, &ord.Customer.ShipCity, &ord.Customer.ShipProvince, &ord.Customer.ShipZipCode, &ord.Customer.ShipPhoneNumber,
		)
		if err != nil {
			return res, err
		}
		res = append(res, ord)
	}
	return res, err
}
