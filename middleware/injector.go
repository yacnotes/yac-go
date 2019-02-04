package middleware

import (
	"github.com/gin-gonic/gin"
	"yac-go/config"
)

func Injector(deps *config.AppDeps) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("deps", deps)
		c.Next()
	}
}
