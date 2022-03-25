package server

import (
	"go-line-bot/lineBotSetting"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

func Callback(ctx *gin.Context) {

	bot := lineBotSetting.GetBot()
	events, err := bot.ParseRequest(ctx.Request)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			ctx.JSON(400, "")
		} else {
			ctx.JSON(500, "")
		}
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch msg := event.Message.(type) {
			case *linebot.TextMessage:
				bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("echo :"+msg.Text))
			}
		}
	}
}
