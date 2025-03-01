package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"shopping-list-backend/internal/database"
	"shopping-list-backend/internal/models"

	"github.com/gin-gonic/gin"
)

// GetProducts responds with the list of all products as JSON.
func GetProducts(c *gin.Context) {
	rows, err := database.DB.Query("SELECT * FROM product")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Title, &product.CategoryID, &product.ListID); err != nil {
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
	err := database.DB.QueryRow("SELECT id, title FROM category WHERE id = ?", newProduct.CategoryID).Scan(&category.ID, &category.Title)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Category not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Check if the list exists
	var list models.List
	err = database.DB.QueryRow("SELECT id, title FROM list WHERE id = ?", newProduct.ListID).Scan(&list.ID, &list.Title)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{"error": "List not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Add the new product to the database
	_, err = database.DB.Exec("INSERT INTO product (title, category_id, list_id) VALUES (?, ?, ?)", newProduct.Title, newProduct.CategoryID, newProduct.ListID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newProduct)
}

// DeteleProduct deletes product by id.
func DeteleProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	_, err = database.DB.Exec("DELETE FROM product WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
