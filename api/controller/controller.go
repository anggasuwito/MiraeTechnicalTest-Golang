package controller

import (
	"MiraeTechnicalTest-Golang/api/controller/c_order"
	"MiraeTechnicalTest-Golang/api/usecase"
	"database/sql"
)

type Controller struct {
	OrderController c_order.OrderControllerInterface
}

func NewController(db *sql.DB) Controller {
	uc := usecase.NewUseCase(db)
	return Controller{
		OrderController: c_order.NewOrderController(uc),
	}
}
