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
	refreshToken, err := query.GetRefreshToken(sendUserId)
	if err != nil {
		log.Print(err)
		Unregistered(bot, event)
		return
	}

	// アクセストークンを取得
	accessToken, err := util.GetAccessToken(refreshToken)
	if err != nil {
		log.Print(err)
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
	imageStr, err := util.ExtractChar(imageData)
	if err != nil {
		log.Print(err)
		Others(bot, event)
		return
	}
	eventData, err := util.FormatJson(imageStr)
	if err != nil {
		log.Print(err)
		Others(bot, event)
		return
	}

	// カレンダーに登録
	err = util.RegisterEvent(eventData, accessToken)
	if err == nil {
		Successful(bot, event, eventData)
	} else if err == util.NoEventError {
		NoEvent(bot, event)
	} else {
		Others(bot, event)
	}

}
