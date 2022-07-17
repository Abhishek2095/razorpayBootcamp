package Routes

import (
	"ecommerce/Handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"user1": "123",
		"user2": "456",
		"user3": "789",
	}))

	authorized.POST("/product", Handlers.AddProduct)

	products := r.Group("/product")
	{
		products.PATCH("", Handlers.UpdateProduct)
		products.GET("", Handlers.GetAllProducts)
		products.GET(":id", Handlers.GetProductById)
	}

	order := r.Group("/order")
	{
		order.GET("/order-history", Handlers.GetOrderHistory)
		order.POST("", Handlers.PlaceOrder)
		order.GET(":id", Handlers.GetOrderById)
		order.GET("/order-history/:id", Handlers.GetCustomerOrderHistory)
		order.POST("multiple", Handlers.PlaceMultipleOrder)
	}

	customer := r.Group("/customer")
	{
		customer.POST("", Handlers.AddCustomer)
		customer.GET(":id", Handlers.GetCustomerById)
	}

	return r
}
