package controlers

import (
	"REST-api/pkg/db"
	"REST-api/pkg/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateBookById(c *gin.Context) {
	id := c.Param("id")

	var book models.Book

	if err := db.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	// Bind the JSON payload to the item struct
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid request body: %v", err)})
		return
	}

	if err := db.DB.Save(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update item"})
		return
	}

	c.JSON(http.StatusOK, book)
}
