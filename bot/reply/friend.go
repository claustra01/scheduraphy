package reply

import (
	"fmt"
	"log"
	"os"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func Friend(bot *linebot.Client, event *linebot.Event) {
	liffUrl := os.Getenv("LIFF_URL")
	message := fmt.Sprintf("友達追加ありがとう!!\nまずはこのリンクからGoogleでログインしてね!!\n%s", liffUrl)
	_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message)).Do()
	if err != nil {
		log.Print(err)
	}
}
