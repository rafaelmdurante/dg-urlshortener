package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetRoutes(g *gin.Engine) {
	g.GET("/health", HealthCheck)

	g.POST("/", CreateShortURL)
	g.GET("/:code", RedirectToTargetURL)
}

func HealthCheck(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}
