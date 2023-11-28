package controlers

import (
	"REST-api/pkg/db"
	"REST-api/pkg/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteBook(c *gin.Context) {

	id := c.Param("id")

	var book models.Book
	if err := db.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	if err := db.DB.Delete(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Item %s deleted", id)})

}
