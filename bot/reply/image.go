package reply

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/claustra01/scheduraphy/token"
	"github.com/claustra01/scheduraphy/util"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func Image(bot *linebot.Client, event *linebot.Event, message *linebot.ImageMessage) {

	// 画像データを取得
	content, err := bot.GetMessageContent(message.ID).Do()
	if err != nil {
		log.Print(err)
		return
	}
	defer content.Content.Close()

	// 画像を一時保存
	filePath := "./image.png"
	file, err := os.Create(filePath)
	if err != nil {
		log.Print(err)
		return
	}
	_, err = io.Copy(file, content.Content)
	if err != nil {
		log.Print(err)
		return
	}
	file.Close()

	// 文字抽出
	file, _ = os.Open(filePath)
	imageStr := util.ExtractChar(file)
	fmt.Println(imageStr)

	sendUserId := event.Source.UserID
	refreshToken := token.GetRefreshToken(sendUserId)

	if refreshToken == "" {
		// ユーザー登録を促すメッセージを返す
		return
	}

	accessToken := token.GetAccessToken(refreshToken)
	fmt.Print(accessToken)

}
