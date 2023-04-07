package main

import (
	"log"
	"os"

	"github.com/LeandroMelloo/go-ecommerce-cart/controllers"
	"github.com/LeandroMelloo/go-ecommerce-cart/database"
	"github.com/LeandroMelloo/go-ecommerce-cart/middleware"
	"github.com/LeandroMelloo/go-ecommerce-cart/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/create/cart", app.CreateCart())
	router.GET("/checkout/cart", app.CheckoutCart())
	router.GET("/delete/item", app.DeleteItem())
	router.GET("/instant/buy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}
