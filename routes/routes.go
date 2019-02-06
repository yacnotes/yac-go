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
	notes.POST("/", actions.PostNewNode)
	notes.GET("/:id", actions.GetNote)
	notes.PATCH("/:id", notImplementedYet)
	notes.DELETE("/:id", actions.DeleteNote)

	entries := api.Group("/entries")
	entries.POST("/:nid", actions.PostNewEntry)
	entries.PATCH("/:nid/:eid", notImplementedYet)
	entries.DELETE("/:nid/:eid", actions.DeleteEntry)

	query := api.Group("/query")
	notesQuery := query.Group("/notes")
	notesQuery.GET("/:year", actions.GetQueryNote)
	notesQuery.GET("/:year/:month", actions.GetQueryNote)
	notesQuery.GET("/:year/:month/:day", actions.GetQueryNote)
}

func notImplementedYet(ctx *gin.Context) {
	ctx.String(http.StatusNotImplemented, "This route is not implemented yet")
}
