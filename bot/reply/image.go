package reply

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/claustra01/scheduraphy/token"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func Image(bot *linebot.Client, event *linebot.Event, message *linebot.ImageMessage) {

	content, err := bot.GetMessageContent(message.ID).Do()
	if err != nil {
		log.Print(err)
		return
	}
	defer content.Content.Close()

	filePath := "./image.png"
	file, err := os.Create(filePath)
	if err != nil {
		log.Print(err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, content.Content)
	if err != nil {
		log.Print(err)
		return
	}

	sendUserId := event.Source.UserID
	refreshToken := token.GetRefreshToken(sendUserId)

	if refreshToken == "" {
		// ユーザー登録を促すメッセージを返す
		return
	}

	accessToken := token.GetAccessToken(refreshToken)
	fmt.Print(accessToken)
}
