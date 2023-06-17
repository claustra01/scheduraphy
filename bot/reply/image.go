package reply

import (
	"fmt"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func Image(bot *linebot.Client, event *linebot.Event) {
	sendUserId := event.Source.UserID
	fmt.Print(sendUserId)
}
