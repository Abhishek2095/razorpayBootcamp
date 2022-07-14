package Handlers

import (
	"ecommerce/Customer"
	"ecommerce/Services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddCustomer(c *gin.Context) {
	var customer Customer.Customer
	c.BindJSON(&customer)

	if err := Services.CustomerServices.CreateCustomer(&customer); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, customer)
	}
}

func GetCustomerById(c *gin.Context) {
	var customer Customer.Customer
	id := c.Params.ByName("id")

	if err := Services.CustomerServices.GetCustomerById(id, &customer); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, customer)
	}
}
