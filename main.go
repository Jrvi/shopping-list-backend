package main

import (
	"shopping-list-backend/internal/database"
	"shopping-list-backend/internal/routes"
)

// Main function to start server
func main() {
	database.InitDB()
	router := routes.SetupRouter()
	router.Run("localhost:8080")
}
