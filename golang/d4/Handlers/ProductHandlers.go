package Handlers

import (
	"ecommerce/Product"
	"ecommerce/Services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllProducts(c *gin.Context) {
	var products []Product.Product

	if err := Services.ProductServices.GetAllProducts(&products); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, products)
	}
}

func GetProductById(c *gin.Context) {
	id := c.Params.ByName("id")
	var product Product.Product

	if err := Services.ProductServices.GetProductById(id, &product); err != nil {
		c.JSON(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func AddProduct(c *gin.Context) {
	var product Product.Product
	c.BindJSON(&product)

	if err := Services.ProductServices.CreateProduct(&product); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func UpdateProduct(c *gin.Context) {
	var product Product.Product
	c.BindJSON(&product)

	if err := Services.ProductServices.UpdateProduct(&product); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}
