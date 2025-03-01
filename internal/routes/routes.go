package routes

import (
	"shopping-list-backend/internal/handlers"
	"shopping-list-backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Public routes (do not require authentication)
	publicRoutes := router.Group("/")
	{
		publicRoutes.POST("/login", handlers.Login)
		publicRoutes.POST("/register", handlers.Register)
	}

	// Protected routes (require authentication)
	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.AuthenticationMiddleware())
	{
		protectedRoutes.GET("/product", handlers.GetProducts)
		protectedRoutes.POST("/product", handlers.PostProduct)
		protectedRoutes.DELETE("/product/:id", handlers.DeteleProduct)
		protectedRoutes.GET("/category", handlers.GetCategories)
		protectedRoutes.POST("/category", handlers.PostCategory)
		protectedRoutes.GET("/list", handlers.GetLists)
		protectedRoutes.GET("/list/:id", handlers.GetList)
		protectedRoutes.POST("list", handlers.PostList)
		protectedRoutes.DELETE("/list/:id", handlers.DeteleList)
		protectedRoutes.PATCH("/list/share", handlers.PatchListShare)
	}

	return router
}
