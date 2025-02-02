package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"shopping-list-backend/internal/database"
	"shopping-list-backend/internal/models"

	"github.com/gin-gonic/gin"
)

// GetLists responds with the list of all shopping-lists as JSON.
func GetLists(c *gin.Context) {
	userID := c.GetInt("user_id")

	rows, err := database.DB.Query("SELECT * FROM lists WHERE user_id = ?", userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
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
	userID := c.GetInt("user_id")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var list models.List
	err = database.DB.QueryRow("SELECT id, title FROM lists WHERE id = ? AND user_id = ?", id, userID).Scan(&list.ID, &list.Title)
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
	userID := c.GetInt("user_id")

	// Call BindJSON to bind the received JSON to newList
	if err := c.BindJSON(&newList); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Add the new list to the database
	_, err := database.DB.Exec("INSERT INTO lists (id, title, user_id) VALUES (?, ?, ?)", newList.ID, newList.Title, userID)
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
