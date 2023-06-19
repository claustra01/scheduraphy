package reply

import (
	"log"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func NoEvent(bot *linebot.Client, event *linebot.Event) {
	message := "予定に関係ない画像が検出されたみたい……\n余計な部分をトリミングするとちゃんと動くかも!"
	_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message)).Do()
	if err != nil {
		log.Print(err)
	}
}
