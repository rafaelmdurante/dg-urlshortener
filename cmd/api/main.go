package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/rafaelmdurante/devgym-urlshortener/internal/database"
	"github.com/rafaelmdurante/devgym-urlshortener/internal/http"
)

func main() {
	ctx := context.Background()

	connectionURI := "postgresql://postgres:postgres@db:5432/urlshortener"
	conn, err := database.NewConnection(ctx, connectionURI)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	g := gin.New()
	g.Use(gin.Recovery())

	http.Configure()
	http.SetRoutes(g)

	err = g.Run(":3000")
	if err != nil {
		return
	}
}
