package routes

import (
	"shopping-list-backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/product", handlers.GetProducts)
	router.POST("/product", handlers.PostProduct)
	router.GET("/categories", handlers.GetCategories)
	router.POST("/categories", handlers.PostCategory)

	return router
}
