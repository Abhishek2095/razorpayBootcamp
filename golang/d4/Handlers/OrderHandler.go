package Handlers

import (
	"ecommerce/DTO"
	"ecommerce/Order"
	"ecommerce/Services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrderHistory(c *gin.Context) {
	panic("implement me")
}

func PlaceOrder(c *gin.Context) {
	var order Order.Order
	c.BindJSON(&order)

	if err := Services.OrderServices.CreateOrder(&order); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

func GetOrderById(c *gin.Context) {
	id := c.Params.ByName("id")
	var order Order.Order

	if err := Services.OrderServices.GetOrderById(id, &order); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

func PlaceMultipleOrder(c *gin.Context) {
	var order DTO.MultipleOrder
	c.BindJSON(&order)

	if err := Services.OrderServices.CreateMultipleOrders(&order); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, order)
	}
}
