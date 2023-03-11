package route

import (
	"MiraeTechnicalTest-Golang/api/controller"
	"github.com/gin-gonic/gin"
)

func Order(r *gin.RouterGroup, c controller.Controller) {
	orderAPI := r.Group("/order")
	orderAPI.GET("/", c.OrderController.GetAllOrder)
}
