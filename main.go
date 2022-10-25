package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prayogayogi/go-restapi-gin/controllers/ordercontroller"
	"github.com/prayogayogi/go-restapi-gin/models"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/orders", ordercontroller.Index)
	r.POST("/orders", ordercontroller.Create)
	r.PUT("/orders/:orderid", ordercontroller.Update)
	r.DELETE("/orders/:orderid", ordercontroller.Delete)
	
	r.Run()
}