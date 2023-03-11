package api

import (
	"MiraeTechnicalTest-Golang/api/controller"
	"MiraeTechnicalTest-Golang/api/route"
	"MiraeTechnicalTest-Golang/config"
	"github.com/gin-gonic/gin"
	"os"
)

func Init() {
	r := gin.Default()
	r.GET("", func(context *gin.Context) { return })
	v1Api := r.Group("/v1")
	c := controller.NewController(config.ConnDB())

	route.Order(v1Api, c)

	r.Run(os.Getenv("ADDRESS"))
}
