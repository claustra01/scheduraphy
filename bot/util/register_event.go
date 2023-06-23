package util

import (
	"context"
	"log"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

func RegisterEvent(data interface{}, accessToken string) string {

	// 画像が読み込めているか確認
	eventType := data.(map[string]interface{})["type"].(string)
	if eventType == "null" {
		return "ImageError"
	}

	// 認証してカレンダーを作成
	config := &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/calendar"},
		Endpoint:     google.Endpoint,
	}

	ctx := context.Background()
	token := &oauth2.Token{
		AccessToken: accessToken,
	}
	client := config.Client(ctx, token)

	srv, err := calendar.New(client)
	if err != nil {
		log.Print(err)
		return ""
	}

	// 登録する予定の情報を作成
	title := data.(map[string]interface{})["title"].(string)
	location := data.(map[string]interface{})["location"].(string)
	startTime := data.(map[string]interface{})["start"].(string)
	endTime := data.(map[string]interface{})["end"].(string)

	event := &calendar.Event{
		Summary:  title,
		Location: location,
		Start: &calendar.EventDateTime{
			DateTime: startTime,
			TimeZone: "Asia/Tokyo",
		},
		End: &calendar.EventDateTime{
			DateTime: endTime,
			TimeZone: "Asia/Tokyo",
		},
	}

	// 予定を登録
	calendarID := "primary"
	event, err = srv.Events.Insert(calendarID, event).Do()
	if err != nil {
		log.Print(err)
		return ""
	}

	return "Successful"
}
