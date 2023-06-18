package reply

import (
	"fmt"
	"log"
	"time"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func Successful(bot *linebot.Client, event *linebot.Event, data interface{}) {

	title := data.(map[string]interface{})["title"].(string)
	location := data.(map[string]interface{})["location"].(string)
	startTime := data.(map[string]interface{})["start"].(string)
	endTime := data.(map[string]interface{})["end"].(string)

	parsedStartTime, _ := time.Parse(time.RFC3339, startTime)
	parsedEndTime, _ := time.Parse(time.RFC3339, endTime)

	message := fmt.Sprintf("予定を登録したよ!!\nタイトル: %s\n会場: %s\n開始: %s月%s日 %s\n終了: %s月%s日 %s",
		title, location,
		parsedStartTime.Format("01"), parsedStartTime.Format("02"), parsedStartTime.Format("15:04"),
		parsedEndTime.Format("01"), parsedEndTime.Format("02"), parsedEndTime.Format("15:04"))

	_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message)).Do()
	if err != nil {
		log.Print(err)
	}

}
