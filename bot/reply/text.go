package reply

import (
	"log"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func Text(bot *linebot.Client, event *linebot.Event, message *linebot.TextMessage) {
	_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do()
	if err != nil {
		log.Print(err)
	}
}
