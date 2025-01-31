package handlers

import (
	"net/http"

	"shopping-list-backend/internal/database"
	"shopping-list-backend/internal/models"

	"github.com/gin-gonic/gin"
)

// GetCategories responds with the list of all categories as JSON.
func GetCategories(c *gin.Context) {
	rows, err := database.DB.Query("SELECT id, title FROM categories")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		if err := rows.Scan(&category.ID, &category.Title); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		categories = append(categories, category)
	}

	c.IndentedJSON(http.StatusOK, categories)
}

// PostCategory adds a category from JSON received in the request body.
func PostCategory(c *gin.Context) {
	var newCategory models.Category

	// Call BindJSON to bind the received JSON to newCategory.
	if err := c.BindJSON(&newCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Add the new category to the database
	_, err := database.DB.Exec("INSERT INTO categories (id, title) VALUES (?, ?)", newCategory.ID, newCategory.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newCategory)
}
