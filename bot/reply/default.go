package reply

import (
	"log"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func Default(bot *linebot.Client, event *linebot.Event) {
	message := `
	画像以外には対応してないんだ!
	ごめんね!!
	`
	_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message)).Do()
	if err != nil {
		log.Print(err)
	}
}
