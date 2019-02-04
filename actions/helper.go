package actions

import (
	"github.com/gin-gonic/gin"
	"yac-go/config"
)

func GetDepsFromCtx(c *gin.Context) *config.AppDeps {
	inf, ok := c.Get("deps")
	if !ok {
		panic("Dependencies have to be injected!")
	}

	deps, ok := inf.(*config.AppDeps)
	if !ok {
		panic("Dependencies object malformed.")
	}

	return deps
}
