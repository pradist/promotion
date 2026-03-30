package router

import (
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

var readBuildInfo = debug.ReadBuildInfo

func AppVersion() string {
	if build, ok := readBuildInfo(); ok && build.Main.Version != "" {
		return build.Main.Version
	}
	return "unknown"
}

func Version(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"version": AppVersion(),
	})
}
