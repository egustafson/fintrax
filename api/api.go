package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitAPI(ctx context.Context, r *gin.Engine) {

	r.GET("/status", statusHdlr)
	r.GET("/live", livelinessHdlr)
	r.GET("/ready", readinessHdlr)
	r.POST("/auth", authHdlr)

	// v0 := r.Group("/v0")

}

func statusHdlr(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}

func livelinessHdlr(c *gin.Context) {
	c.JSON(http.StatusOK, "alive")
}

func readinessHdlr(c *gin.Context) {
	c.JSON(http.StatusOK, "ready")
}
