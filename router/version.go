package router

import (
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func AppVersion() string {
	if build, ok := debug.ReadBuildInfo(); ok {
		return build.Main.Version
	}
	return "unknown"
}

func Version(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"version": AppVersion(),
	})
}
