package controllers

import (
	"challenge-3-chapter-2/database"
	"challenge-3-chapter-2/models"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBookAll(ctx *gin.Context) {
	var bookDatas []models.Book

	getBooks, err := database.GetBookAllDB(bookDatas)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Gagal mendapatkan request")
		return
	}

	if getBooks == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"data": []string{},
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": getBooks,
	})
}

func GetBookId(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	var bookDatas models.Book

	convID, err := strconv.Atoi(bookID)
	if err != nil {
		log.Println("Gagal Mengconvert")
		return
	}

	book, err := database.GetBookIdDB(convID, bookDatas)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("book with id %v not found", convID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func CreateBook(ctx *gin.Context) {
	var bookDatas models.Book

	if err := ctx.ShouldBindJSON(&bookDatas); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	_, err := database.CreateBookDB(bookDatas)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Gagal membuat buku")
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Book has created",
	})
}

func UpdateBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	var bookDatas models.Book

	convID, err := strconv.Atoi(bookID)
	if err != nil {
		log.Println("Gagal Mengconvert")
		return
	}

	if err = ctx.ShouldBindJSON(&bookDatas); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	_, err = database.UpdateBookDB(convID, bookDatas)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("book with id %d not found", convID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("book with id %d succesfully updated", convID),
	})
}

func DeleteBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")

	convID, err := strconv.Atoi(bookID)
	if err != nil {
		log.Println("Gagal Mengconvert")
		return
	}

	err = database.DeleteBookDB(convID)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("book with id %d not found", convID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("book with id %d succesfully deleted", convID),
	})
}
