package uc_order

import (
	"MiraeTechnicalTest-Golang/api/models"
	"MiraeTechnicalTest-Golang/api/repository"
	"MiraeTechnicalTest-Golang/helper/pagination"
)

type OrderUC struct {
	Repo repository.Repository
}

type OrderUCInterface interface {
	GetAllOrderUC(req pagination.Request) (res []models.OrderResponse, meta pagination.Response, err error)
}

func NewOrderUC(repo repository.Repository) OrderUCInterface {
	return &OrderUC{Repo: repo}
}
