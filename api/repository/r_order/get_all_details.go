package r_order

import "MiraeTechnicalTest-Golang/api/models"

func (r OrderRepo) GetAllOrderDetailsRepo(orderID int) (res []models.OrderDetailsResponse, err error) {
	rows, err := r.DB.
		Query(`
			SELECT 
			    od.order_detail_id,od.quantity,od.unit_price,od.discount, 
			    p.product_id,p.product_name,p.unit_price,p.in_stock
			FROM order_details od
			LEFT JOIN products p ON p.product_id = od.product_id
			WHERE
			    od.order_id = ?
		`, orderID)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var ord models.OrderDetailsResponse
		err = rows.Scan(
			&ord.OrderDetailID, &ord.Quantity, &ord.UnitPrice, &ord.Discount,
			&ord.Product.ProductID, &ord.Product.ProductName, &ord.Product.UnitPrice, &ord.Product.InStock,
		)
		if err != nil {
			return res, err
		}
		res = append(res, ord)
	}
	return res, err
}
