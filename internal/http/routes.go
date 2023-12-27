package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetRoutes(g *gin.Engine) {
	g.GET("/health", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})

	g.POST("/", CreateShortURL)
	g.GET("/:code", RedirectToTargetURL)
}
