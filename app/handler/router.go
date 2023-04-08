package handler

import (
	"jwt-go/app/middleware"
	"jwt-go/app/services"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", services.UserRegister)
		userRouter.POST("/login", services.UserLogin)
	}

	productRouter := router.Group("/products")
	{
		productRouter.Use(middleware.Authentication())
		productRouter.POST("/", services.CreateProduct)
		productRouter.GET("/:productId", middleware.ProductAuthorization(), services.ReadProduct)
		productRouter.PUT("/:productId", middleware.ProductAuthorization(), services.UpdateProduct)
		productRouter.DELETE("/:productId", middleware.ProductAuthorization(), services.DeleteProduct)
	}
	return router
}
