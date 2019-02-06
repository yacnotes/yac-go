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

	if full {
		notes, err := service.GetAll(db)
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
		ctx.JSON(http.StatusPreconditionFailed, gin.H{"error": "A note for this day already exists", "id": otherId})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": id, "note": n})
}

func GetNote(ctx *gin.Context) {
	strId := ctx.Param("id")
	nid, err := strconv.Atoi(strId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := GetDepsFromCtx(ctx).Db

	n, err, _ := service.GetByIdentifier(db, nid)
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

	notes, err := service.Query(db, year, month, day)
	if err != nil {
		log.Panic("Failed to query for notes:", err)
	}

	ctx.JSON(http.StatusOK, notes)
}

func DeleteNote(ctx *gin.Context) {
	db := GetDepsFromCtx(ctx).Db

	identifier, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = service.DeleteByIdentifier(db, identifier)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
