package reply

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/claustra01/scheduraphy/query"
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

	// 画像データをバイトスライスに読み込む
	imageData, err := ioutil.ReadAll(content.Content)
	if err != nil {
		log.Print(err)
		return
	}

	// 文字抽出
	imageStr := util.ExtractChar(imageData)

	// 抽出結果を返信（後で消す）
	_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(imageStr)).Do()
	if err != nil {
		log.Print(err)
	}

	sendUserId := event.Source.UserID
	refreshToken := query.GetRefreshToken(sendUserId)

	if refreshToken == "" {
		// ユーザー登録を促すメッセージを返す
		return
	}

	accessToken := util.GetAccessToken(refreshToken)
	fmt.Print(accessToken)

}
