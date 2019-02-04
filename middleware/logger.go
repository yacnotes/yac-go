package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
	"yac-go/config"
	"yac-go/log"
)

func Logger(deps *config.AppDeps) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Next()
		latency := time.Since(t)
		log.Debug("Request latency:", latency)
	}
}
