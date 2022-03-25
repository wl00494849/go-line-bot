package lineBotSetting

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func CreateBot(sectet string, token string) {

	godotenv.Load()

	b, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_ACCESS_TOKEN"),
	)

	if err != nil {
		panic(err)
	}

	bot = b
}

func GetBot() *linebot.Client {
	return bot
}
