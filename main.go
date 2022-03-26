package main

import (
	"flag"
	"go-line-bot/lineBotSetting"
	"go-line-bot/server"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	var port string
	flag.StringVar(&port, "p", os.Getenv("Port"), "port")

	lineBotSetting.CreateBot()

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
