package routes

import (
	"shopping-list-backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/product", handlers.GetProducts)
	router.POST("/product", handlers.PostProduct)
	router.DELETE("/product/:id", handlers.DeteleProduct)
	router.GET("/category", handlers.GetCategories)
	router.POST("/category", handlers.PostCategory)
	router.GET("/list", handlers.GetLists)
	router.GET("/list/:id", handlers.GetList)
	router.POST("list", handlers.PostList)
	router.DELETE("/list/:id", handlers.DeteleList)

	return router
}
