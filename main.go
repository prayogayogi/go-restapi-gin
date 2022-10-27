package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prayogayogi/go-restapi-gin/controllers/ordercontroller"
	"github.com/prayogayogi/go-restapi-gin/models"
)

func main() {
	route := gin.Default()
	models.ConnectDatabase()

	route.GET("/orders", ordercontroller.Index)
	route.POST("/orders", ordercontroller.Create)
	route.PUT("/orders/:orderid", ordercontroller.Update)
	route.DELETE("/orders/:orderid", ordercontroller.Delete)
	
	route.Run()
}