package main

import (
	"flag"
	"go-line-bot/server"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	var port string
	flag.StringVar(&port, "p", ":6666", "port")

	if p := os.Getenv("PORT"); len(p) > 0 {
		port = ":" + p
	}

	app := gin.Default()

	app.POST("/callback", server.Callback)
	app.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]string{"test": "success"})
	})

	err := app.Run(port)

	if err != nil {
		panic(err)
	}
}
