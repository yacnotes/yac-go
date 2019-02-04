package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yac-go/actions"
	"yac-go/config"
)

func LoadRoutes(r *gin.Engine, deps *config.AppDeps) {
	api := r.Group("/api/v1")
	api.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, deps.Config.AppInfo)
	})

	notes := api.Group("/notes")
	notes.GET("/", actions.GetAllNotes)
	notes.GET("/:id", notImplementedYet)
	notes.POST("/:id", notImplementedYet)
}

func notImplementedYet(ctx *gin.Context) {
	ctx.String(http.StatusNotImplemented, "This route is not implemented yet")
}
