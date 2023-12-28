package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/rafaelmdurante/devgym-urlshortener/internal"
	"github.com/rafaelmdurante/devgym-urlshortener/internal/database"
	"github.com/rafaelmdurante/devgym-urlshortener/internal/url"
	"net/http"
	"time"
)

var service url.Service

func GetErrorMessage(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func Configure() {
	service = url.Service{
		Repository: &url.RepositoryPostgres{Conn: database.Connection},
	}
}

func CreateShortURL(ctx *gin.Context) {
	var shortURL internal.URL

	if err := ctx.BindJSON(&shortURL); err != nil {
		ctx.JSON(http.StatusBadRequest, GetErrorMessage(err))

		return
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	response, err := service.Create(ctxTimeout, shortURL)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, GetErrorMessage(err))

		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func RedirectToTargetURL(ctx *gin.Context) {
	code := ctx.Param("code")

	if code == "" {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": url.ErrInvalidURLCode,
		})
	}

	// get shortened url
	url, err := service.FindOneByCode(ctx, code)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)

		return
	}

	ctx.Redirect(http.StatusMovedPermanently, url.TargetURL)
}
