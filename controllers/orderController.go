package orderController

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mcsans/assignment2-019/models"
)

func Index(c *gin.Context) {
	var orders []models.Order

	models.DB.Preload("Items").Find(&orders)
	c.JSON(http.StatusOK, gin.H{"orders" : orders})
}

func Show(c *gin.Context) {
}

func Create(c *gin.Context) {
	var order models.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		// c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{"message": errorMessages})
		return
	}

	models.DB.Create(&order)
	c.JSON(http.StatusOK, gin.H{"order" : order})
}

func Update(c *gin.Context) {
}

func Delete(c *gin.Context) {
}