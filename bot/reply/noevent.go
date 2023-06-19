package reply

import (
	"log"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func NoEvent(bot *linebot.Client, event *linebot.Event) {
	message := "これは予定に関係ない画像だと思うんだ……\n余計な部分をトリミングしてからもう一回試してみて!"
	_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message)).Do()
	if err != nil {
		log.Print(err)
	}
}
