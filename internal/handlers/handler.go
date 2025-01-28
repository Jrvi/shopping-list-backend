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
	rows, err := database.DB.Query("SELECT * FROM products")
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
	err := database.DB.QueryRow("SELECT id, title FROM categories WHERE id = ?", newProduct.CategoryID).Scan(&category.ID, &category.Title)
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
	err = database.DB.QueryRow("SELECT id, title FROM lists WHERE id = ?", newProduct.ListID).Scan(&list.ID, &list.Title)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{"error": "List not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Add the new product to the database
	_, err = database.DB.Exec("INSERT INTO products (id, title, category_id, list_id) VALUES (?, ?, ?, ?)", newProduct.ID, newProduct.Title, newProduct.CategoryID, newProduct.ListID)
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
	_, err = database.DB.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
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

// GetLists responds with the list of all shopping-lists as JSON.
func GetLists(c *gin.Context) {
	rows, err := database.DB.Query("SELECT * FROM lists")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	defer rows.Close()

	var lists []models.List
	for rows.Next() {
		var list models.List
		if err := rows.Scan(&list.ID, &list.Title); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		lists = append(lists, list)
	}

	c.IndentedJSON(http.StatusOK, lists)
}

// GetList responds with the shopping-list as JSON.
func GetList(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var list models.List
	err = database.DB.QueryRow("SELECT id, title FROM lists WHERE id = ?", id).Scan(&list.ID, &list.Title)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "List not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, list)
}

// PostList adds a list from JSON received in the request body.
func PostList(c *gin.Context) {
	var newList models.List

	// Call BindJSON to bind the received JSON to newList
	if err := c.BindJSON(&newList); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Add the new list to the database
	_, err := database.DB.Exec("INSERT INTO lists (id, title) VALUES (?, ?)", newList.ID, newList.Title)
	if err != nil {
		c.JSON((http.StatusInternalServerError), gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, newList)
}

// DeteleList deletes product by id.
func DeteleList(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	_, err = database.DB.Exec("DELETE FROM lists WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
