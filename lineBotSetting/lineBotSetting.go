package lineBotSetting

import (
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func init() {

	os.Setenv("CHANNEL_SECRET", "")
	os.Setenv("CHANNEL_ACCESS_TOKEN", "")

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
