package actions

import (
	"github.com/gin-gonic/gin"
	"yac-go/service/note"
)

func GetAllNotes(ctx *gin.Context) {
	db := GetDepsFromCtx(ctx).Db
	ctx.JSON(200, note.GetAll(db, 0, 30))
}
