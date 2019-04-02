package actions

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"yac-go/log"
	"yac-go/model/book"
	service "yac-go/service/book"
	"yac-go/service/note"
)

func PostNewBook(ctx *gin.Context) {
	db := GetDepsFromCtx(ctx).Db

	b := &book.Book{}
	if err := ctx.ShouldBindJSON(b); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := service.Add(db, b)
	if err != nil {
		log.Panic("Insert error:", err)
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": id, "book": b})
}

func GetAllBooks(ctx *gin.Context) {
	db := GetDepsFromCtx(ctx).Db

	books, err := service.GetAll(db)
	if err != nil {
		log.Panic("Error while loading all books:", err)
	}
	ctx.JSON(http.StatusOK, books)
	return
}

func DeleteBook(ctx *gin.Context) {
	db := GetDepsFromCtx(ctx).Db

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// delete book
	if err := service.Delete(db, id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// delete notes from book
	// this does not happen in book server because it will cause an import cycle
	notes, err := note.GetAll(db, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, n := range notes {
		if err := note.DeleteById(db, n.Id); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	ctx.Status(http.StatusOK)
}
