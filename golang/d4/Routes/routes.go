package Routes

import (
	"ecommerce/Handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	products := r.Group("/product")
	{
		products.POST("", Handlers.AddProduct)
		products.PATCH(":id", Handlers.UpdateProduct)
		products.GET("", Handlers.GetAllProducts)
		products.GET(":id", Handlers.GetProductById)
	}

	order := r.Group("/order")
	{
		order.GET("/order-history", Handlers.GetOrderHistory)
		order.POST("", Handlers.PlaceOrder)
		order.GET(":id", Handlers.GetOrderById)
	}

	customer := r.Group("/customer")
	{
		customer.POST("", Handlers.AddCustomer)
		customer.GET(":id", Handlers.GetCustomerById)
	}

	return r
}
