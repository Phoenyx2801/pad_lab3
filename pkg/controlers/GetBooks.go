package controlers

import (
	"REST-api/pkg/db"
	"REST-api/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetBooks(ctx *gin.Context) {

	id := ctx.Param("id")

	var books models.Book
	db.DB.Find(&books, id)

	ctx.Set("responseBody", books)
	ctx.JSON(http.StatusOK, books)
}
