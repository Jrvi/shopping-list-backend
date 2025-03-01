package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"shopping-list-backend/internal/database"
	"shopping-list-backend/internal/models"

	"github.com/gin-gonic/gin"
)

// GetLists responds with the list of all shopping-lists as JSON.
func GetLists(c *gin.Context) {
	userID, exists := c.Get("user_id")
	fmt.Println(userID)
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unauthorized"})
		return
	}

	query := `
	SELECT l.id, l.title, l.user_id, 
       CASE WHEN sl.list_id IS NOT NULL THEN 1 ELSE 0 END AS is_shared
	FROM list l
	LEFT JOIN list_shares sl 
    	ON l.id = sl.list_id AND sl.shared_with_user_id = ?
	WHERE l.user_id = ? OR sl.shared_with_user_id = ?
`

	rows, err := database.DB.Query(query, userID, userID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var lists []models.List
	for rows.Next() {
		var list models.List
		var isShared bool

		if err := rows.Scan(&list.ID, &list.Title, &list.UserId, &isShared); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		list.IsShared = isShared
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
	err = database.DB.QueryRow("SELECT id, title FROM list WHERE id = ? AND user_id = ?", id, userID).Scan(&list.ID, &list.Title)
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
	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Call BindJSON to bind the received JSON to newList
	if err := c.BindJSON(&newList); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Add the new list to the database
	_, err := database.DB.Exec("INSERT INTO list (title, user_id) VALUES (?, ?)", newList.Title, userID)
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
	_, err = database.DB.Exec("DELETE FROM list WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

// PatchListShare shares list with user
func PatchListShare(c *gin.Context) {
	var newShared models.SharedList
	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if err := c.BindJSON(&newShared); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	var ownedByUser bool
	err := database.DB.QueryRow(`SELECT  EXISTS(SELECT 1 FROM list WHERE id = ? AND user_id = ?)`, newShared.ListId, userID).Scan(&ownedByUser)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if !ownedByUser {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't own this list"})
		return
	}

	_, err = database.DB.Exec(`INSERT INTO list_shares (list_id, shared_with_user_id) VALUES (?, ?)`, newShared.ListId, newShared.SharedWithUserId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "List shared successfully"})
}
