package reply

import (
	"log"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func Others(bot *linebot.Client, event *linebot.Event) {
	message := "原因不明のエラーが起きちゃったみたい……\nもう一回試して欲しいな!"
	_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message)).Do()
	if err != nil {
		log.Print(err)
	}
}
