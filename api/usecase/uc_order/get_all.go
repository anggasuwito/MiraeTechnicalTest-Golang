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

	total, err := uc.Repo.OrderRepo.CountTotalOrderRepo()
	if err != nil {
		return res, meta, err
	}
	meta = pagination.GetPaginationResponse(req.Page, limit, total)
	return res, meta, err
}
