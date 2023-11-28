package controlers

import (
	"REST-api/pkg/db"
	"REST-api/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddBook(ctx *gin.Context) {
	var input models.AddBook
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{Title: input.Title, Author: input.Title}
	db.DB.Create(&book)

	ctx.JSON(http.StatusOK, gin.H{"data": book})
}
