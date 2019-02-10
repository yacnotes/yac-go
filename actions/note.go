package actions

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"yac-go/log"
	"yac-go/model/note"
	service "yac-go/service/note"
)

func GetAllNotes(ctx *gin.Context) {
	db := GetDepsFromCtx(ctx).Db

	_, full := ctx.GetQuery("full")
	rawBookId, ok := ctx.GetQuery("book")
	bookId := 0
	if ok {
		var err error
		bookId, err = strconv.Atoi(rawBookId)
		if err != nil {
			log.Panic("Error while converting string to bookId")
		}
	}

	if full || bookId != 0 {
		notes, err := service.GetAll(db, bookId)
		if err != nil {
			log.Panic("Error while loading all notes:", err)
		}
		ctx.JSON(http.StatusOK, notes)
		return
	}

	notes, err := service.GetAllIds(db)
	if err != nil {
		log.Panic("Error while loading all notes:", err)
	}
	ctx.JSON(http.StatusOK, notes)
}

func PostNewNode(ctx *gin.Context) {
	db := GetDepsFromCtx(ctx).Db

	n := &note.Note{}
	if err := ctx.ShouldBindJSON(n); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, otherId, err := service.Add(db, n)
	if id == 0 && otherId == 0 && err != nil {
		log.Panic("Insert error:", err)
	}

	if id == 0 && otherId != 0 {
		ctx.JSON(http.StatusPreconditionFailed, gin.H{"error": err.Error(), "id": otherId})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": id, "note": n})
}

func GetNote(ctx *gin.Context) {
	nid, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := GetDepsFromCtx(ctx).Db

	n, err := service.GetById(db, nid)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, n)
}

func GetQueryNote(ctx *gin.Context) {
	db := GetDepsFromCtx(ctx).Db

	rawDay := ctx.Param("day")
	rawMonth := ctx.Param("month")
	rawYear := ctx.Param("year")

	bookId, _ := strconv.Atoi(ctx.Query("book"))

	day, err := strconv.Atoi(rawDay)
	if err != nil {
		day = 0
	}
	month, err := strconv.Atoi(rawMonth)
	if err != nil {
		month = 0
	}
	year, err := strconv.Atoi(rawYear)
	if err != nil {
		year = 0
	}

	notes, err := service.Query(db, year, month, day, bookId)
	if err != nil {
		log.Panic("Failed to query for notes:", err)
	}

	ctx.JSON(http.StatusOK, notes)
}

func DeleteNote(ctx *gin.Context) {
	db := GetDepsFromCtx(ctx).Db

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = service.DeleteById(db, id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
