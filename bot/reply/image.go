package reply

import (
	"fmt"

	"github.com/claustra01/scheduraphy/token"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func Image(bot *linebot.Client, event *linebot.Event) {

	sendUserId := event.Source.UserID
	refreshToken := token.GetRefreshToken(sendUserId)

	if refreshToken == "" {
		// ユーザー登録を促すメッセージを返す
		return
	}

	accessToken := token.GetAccessToken(refreshToken)
	fmt.Print(accessToken)
}
