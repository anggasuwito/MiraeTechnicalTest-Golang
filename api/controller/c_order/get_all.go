package c_order

import (
	"MiraeTechnicalTest-Golang/helper/pagination"
	"MiraeTechnicalTest-Golang/helper/response"
	"github.com/gin-gonic/gin"
)

func (c OrderController) GetAllOrder(g *gin.Context) {
	var req pagination.Request
	g.ShouldBind(&req)
	res, meta, err := c.UC.OrderUC.GetAllOrderUC(req)
	response.JSON(response.Response{
		Ctx:   g,
		Data:  res,
		Meta:  meta,
		Error: err,
	})
}
