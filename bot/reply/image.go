package reply

import (
	"io/ioutil"
	"log"

	"github.com/claustra01/scheduraphy/query"
	"github.com/claustra01/scheduraphy/util"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func Image(bot *linebot.Client, event *linebot.Event, message *linebot.ImageMessage) {

	// DBからリフレッシュトークンを取得
	sendUserId := event.Source.UserID
	refreshToken := query.GetRefreshToken(sendUserId)
	if refreshToken == "" {
		Unregistered(bot, event)
		return
	}

	// アクセストークンを取得
	accessToken := util.GetAccessToken(refreshToken)
	if accessToken == "" {
		Expired(bot, event)
		return
	}

	// 送信された画像データを取得
	content, err := bot.GetMessageContent(message.ID).Do()
	if err != nil {
		log.Print(err)
		return
	}
	defer content.Content.Close()

	// 画像データをバイト配列に読み込む
	imageData, err := ioutil.ReadAll(content.Content)
	if err != nil {
		log.Print(err)
		return
	}

	// 文字抽出&整形
	imageStr := util.ExtractChar(imageData)
	eventData := util.FormatJson(imageStr)

	// カレンダーに登録
	result := util.RegisterEvent(eventData, accessToken)
	switch result {
	case "Successful":
		Successful(bot, event, eventData)
	case "ImageError":
		NoEvent(bot, event)
	default:
		Others(bot, event)
	}

}
