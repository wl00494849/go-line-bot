package server

import (
	"encoding/json"
	"go-line-bot/lineBotSetting"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

type foodList struct {
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
					getList(event, bot)
				case "吃什麼":
					eatWhat(event, bot)
				case "刪除":
				default:
					if msg.Text[:7] == "新增：" {
						additem(event, bot, msg.Text[7:])
					}
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

func getList(event *linebot.Event, bot *linebot.Client) {
	list := &[]foodList{}
	jsonfile, _ := os.Open("list.json")
	defer jsonfile.Close()

	bytesValue, _ := ioutil.ReadAll(jsonfile)
	json.Unmarshal(bytesValue, list)

	str := "清單: \n"

	for index, v := range *list {
		str += strconv.Itoa(index+1) + ". " + v.Name + "\n"
	}

	bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(str)).Do()
}

func eatWhat(event *linebot.Event, bot *linebot.Client) {
	list := []foodList{}
	jsonfile, _ := os.Open("list.json")
	defer jsonfile.Close()

	bytesValue, _ := ioutil.ReadAll(jsonfile)
	json.Unmarshal(bytesValue, &list)

	i := rand.Intn(len(list))

	bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(list[i].Name)).Do()
}

func additem(event *linebot.Event, bot *linebot.Client, item string) {
	list := []foodList{}
	jsonfile, _ := os.Open("list.json")
	defer jsonfile.Close()

	bytesValue, _ := ioutil.ReadAll(jsonfile)
	json.Unmarshal(bytesValue, &list)
	list = append(list, foodList{Name: item})

	data, _ := json.Marshal(list)
	ioutil.WriteFile("list.json", data, 0644)

	bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("success")).Do()
}
