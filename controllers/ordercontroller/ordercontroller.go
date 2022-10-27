package ordercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prayogayogi/go-restapi-gin/models"
)
func Index(c *gin.Context){
	var orders []models.Order

	models.DB.Find(&orders)
	c.JSON(http.StatusOK, orders)
}

func Create(c *gin.Context){
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	models.DB.Create(&order)
	c.JSON(http.StatusOK, gin.H{"order": order})
}

func Update(c *gin.Context){
	var order models.Order
	id := c.Param("orderid")

	if err := c.ShouldBindJSON(&order); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&order).Where("order_id = ?", id).Updates(&order).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})
}

func Delete(c *gin.Context){
	var order models.Order
	id := c.Param("orderid")

	if models.DB.Model(&order).Where("order_id = ?", id).Delete(&order).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat di hapus"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}