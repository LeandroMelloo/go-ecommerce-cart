package routes

import (
	"github.com/LeandroMelloo/go-ecommerce-cart/controllers"
	"github.com/gin-gonic/gin"
)

func userRoute(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users/search", controllers.UserSearch())
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/product/create", controllers.ProductCreate())
	incomingRoutes.GET("/product/view", controllers.ProductView())
}
