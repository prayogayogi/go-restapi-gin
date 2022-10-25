package ordercontroller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/prayogayogi/go-restapi-gin/models"
)
func Index(c *gin.Context){
	var orders models.Order

	models.DB.Find(&orders)
	c.JSON(http.StatusOK, orders)
}

func Create(c *gin.Context){
	var order models.Order

	if err := c.ShouldBindJSON(&order); err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return

		models.DB.Create(&order)
		c.JSON(http.StatusOK, gin.H{"order":order})
	}
}

func Update(c *gin.Context){

}

func Delete(c *gin.Context){

}