package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/rafaelmdurante/devgym-urlshortener/internal"
	"github.com/rafaelmdurante/devgym-urlshortener/internal/database"
	"github.com/rafaelmdurante/devgym-urlshortener/internal/shorturl"
	"net/http"
	"time"
)

var service shorturl.Service

func GetErrorMessage(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func Configure() {
	service = shorturl.Service{
		Repository: &shorturl.RepositoryPostgres{Conn: database.Connection},
	}
}

func CreateShortURL(ctx *gin.Context) {
	var shortURL internal.ShortenedURL

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