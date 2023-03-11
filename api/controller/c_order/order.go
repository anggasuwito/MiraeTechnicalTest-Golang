package c_order

import (
	"MiraeTechnicalTest-Golang/api/usecase"
	"github.com/gin-gonic/gin"
)

type OrderController struct {
	UC usecase.UseCase
}

type OrderControllerInterface interface {
	GetAllOrder(g *gin.Context)
}

func NewOrderController(uc usecase.UseCase) OrderControllerInterface {
	return &OrderController{UC: uc}
}
