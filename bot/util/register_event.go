package util

import (
	"context"
	"fmt"
	"log"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

func RegisterEvent(data interface{}, accessToken string) {

	eventType := data.(map[string]interface{})["type"].(string)
	if eventType == "null" {

		// 画像が読み込めなかった返信を行う
		log.Print("aaaaaaaaaaaaaaaaaaaaaa")
		return
	}

	title := data.(map[string]interface{})["title"].(string)
	location := data.(map[string]interface{})["location"].(string)
	startTime := data.(map[string]interface{})["start"].(string)
	endTime := data.(map[string]interface{})["end"].(string)

	config := &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/calendar"},
		Endpoint:     google.Endpoint,
	}

	log.Print(title, location, startTime, endTime)

	ctx := context.Background()
	token := &oauth2.Token{
		AccessToken: accessToken,
	}
	client := config.Client(ctx, token)

	srv, err := calendar.New(client)
	if err != nil {
		log.Fatal("クライアントの作成に失敗しました:", err)
	}

	// 登録する予定の情報を作成します
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

	// 予定を登録します
	calendarID := "primary" // 予定を登録するカレンダーのID（"primary"はデフォルトのカレンダー）
	event, err = srv.Events.Insert(calendarID, event).Do()
	if err != nil {
		log.Fatal("予定の登録に失敗しました:", err)
	}

	// 取得した予定の詳細を表示します
	fmt.Println("登録した予定の詳細:")
	fmt.Println("タイトル:", event.Summary)
	fmt.Println("場所:", event.Location)
	fmt.Println("説明:", event.Description)
	fmt.Println("開始日時:", event.Start.DateTime)
	fmt.Println("終了日時:", event.End.DateTime)
}
