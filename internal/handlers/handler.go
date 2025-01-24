package handlers

import (
	"database/sql"
	"net/http"

	"shopping-list-backend/internal/database"
	"shopping-list-backend/internal/models"

	"github.com/gin-gonic/gin"
)

// GetProducts responds with the list of all products as JSON.
func GetProducts(c *gin.Context) {
	rows, err := database.DB.Query("SELECT * FROM products")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Title, &product.CategoryID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}
		products = append(products, product)
	}

	c.IndentedJSON(http.StatusOK, products)
}

// PostProduct adds a product from JSON received in the request body.
func PostProduct(c *gin.Context) {
	var newProduct models.Product

	// Call BindJSON to bind the received JSON to newProduct.
	if err := c.BindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Check if the category exists
	var category models.Category
	err := database.DB.QueryRow("SELECT id, title FROM categories WHERE id = ?", newProduct.CategoryID).Scan(&category.ID, &category.Title)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Category not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Add the new product to the database
	_, err = database.DB.Exec("INSERT INTO products (id, title, category_id) VALUES (?, ?, ?)", newProduct.ID, newProduct.Title, newProduct.CategoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newProduct)
}

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
