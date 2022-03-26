package lineBotSetting

import (
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func init() {

	os.Setenv("CHANNEL_SECRET", "f46e97eb2b29c171339b263b64019371")
	os.Setenv("CHANNEL_ACCESS_TOKEN", "DxZ2LdFclcvzOTFz0HjD07Mk9+lj3Uw5jOF/gBY1reZaTopY3ppLJk1bImG0e9b4iYGdyzqVgSWE3KlLzojbYON3KkNzH9/ztFINaWpGMQ0C5bwKMQPyfSnBYu/FYV3N8fDeblBgvj0F9CcQ1Lp3hQdB04t89/1O/w1cDnyilFU=")

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
