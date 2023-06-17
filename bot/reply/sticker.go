package reply

import (
	"fmt"
	"log"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func Sticker(bot *linebot.Client, event *linebot.Event, message *linebot.StickerMessage) {
	replyMessage := fmt.Sprintf(
		"sticker id is %s, stickerResourceType is %s", message.StickerID, message.StickerResourceType)
	_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do()
	if err != nil {
		log.Print(err)
	}
}
