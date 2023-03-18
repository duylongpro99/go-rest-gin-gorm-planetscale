package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"product-listing-api/controller"
	"product-listing-api/database"
	"product-listing-api/model"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	runApp()
}

func loadEnv() {
	if err := godotenv.Load(".env.local"); err != nil {
		handleError(errors.New("error loading .env file"))
	}
}

func loadDatabase() {
	if err := database.Connect(); err != nil {
		handleError(err)
	}

	if err := database.Database.AutoMigrate(&model.Product{}); err != nil {
		handleError(err)
	}

	fmt.Println(("Migrate executed successfully"))
}

func handleError(err error) {
	log.Fatal(err)
}

func runApp() {
	router := gin.Default()

	userRoutes := router.Group("/product")
	userRoutes.POST("", controller.CreateProduct)
	userRoutes.GET("", controller.GetAllProducts)
	userRoutes.GET("/:id", controller.GetProduct)
	userRoutes.PUT("/:id", controller.UpdateProduct)
	userRoutes.DELETE("/:id", controller.DeleteProduct)

	if err := router.Run(os.Getenv("PORT")); err != nil {
		handleError(err)
	}

	fmt.Println("Server running on port 8000")
}
