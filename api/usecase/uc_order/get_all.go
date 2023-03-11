package uc_order

import (
	"MiraeTechnicalTest-Golang/api/models"
	"MiraeTechnicalTest-Golang/helper/pagination"
)

func (uc OrderUC) GetAllOrderUC(req pagination.Request) (res []models.OrderResponse, meta pagination.Response, err error) {
	offset, limit := pagination.GetPagination(req)
	res, err = uc.Repo.OrderRepo.GetAllOrderRepo(offset, limit)
	if err != nil {
		return res, meta, err
	}

	for i := range res {
		res[i].Details, err = uc.Repo.OrderRepo.GetAllOrderDetailsRepo(res[i].OrderID)
		if err != nil {
			return res, meta, err
		}
		for i2 := range res[i].Details {
			unitPrice := res[i].Details[i2].UnitPrice
			discount := res[i].Details[i2].Discount
			detailPrice := unitPrice - (unitPrice * discount / 100)
			res[i].Details[i2].Price = detailPrice
			res[i].TotalPrice += detailPrice
		}

		totalPrice := res[i].TotalPrice
		taxes := res[i].Taxes
		priceWithTaxes := totalPrice * taxes / 100
		freightCharge := res[i].FreightCharge
		res[i].FinalTotalPrice = totalPrice + priceWithTaxes + freightCharge
	}

	total, err := uc.Repo.OrderRepo.CountTotalOrderRepo()
	if err != nil {
		return res, meta, err
	}
	meta = pagination.GetPaginationResponse(req.Page, limit, total)
	return res, meta, err
}
