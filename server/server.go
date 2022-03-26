package server

import (
	"encoding/json"
	"go-line-bot/lineBotSetting"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

type foodList struct {
	Id   int
	Name string
}

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
				switch msg.Text {
				case "測試":
					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("目前連線正常")).Do()
				case "清單":
					list := &[]foodList{}
					jsonfile, _ := os.Open("list.json")
					defer jsonfile.Close()

					bytesValue, _ := ioutil.ReadAll(jsonfile)
					json.Unmarshal(bytesValue, list)

					str := "清單: %0D%0A"

					for _, v := range *list {
						str += strconv.Itoa(v.Id) + ". " + v.Name + "%0D%0A"
					}

					bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(str)).Do()

				case "新增梗圖":
				case "刪除梗圖":

				}
			}
		}
	}
}

func JsonFileTest(ctx *gin.Context) {
	list := &[]foodList{}
	jsonfile, err := os.Open("list.json")

	if err != nil {
		panic(err)
	}

	defer jsonfile.Close()

	bytesValue, _ := ioutil.ReadAll(jsonfile)
	json.Unmarshal(bytesValue, &list)

	ctx.JSON(200, list)
}
