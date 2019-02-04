package actions

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
