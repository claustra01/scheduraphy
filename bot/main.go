package main

import (
	"log"
	"net/http"
	"os"

	"github.com/claustra01/scheduraphy/db"
	"github.com/claustra01/scheduraphy/query"
	"github.com/claustra01/scheduraphy/reply"
	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("[ERROR] loading: %v", err)
	}

	// migrate DB
	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		db.Migrate()
		return
	}

	bot, err := linebot.New(
		os.Getenv("LINE_BOT_CHANNEL_SECRET"),
		os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Print("[INFO] Bot is running...")
		db.Connect()
	}

	// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		log.Print(err)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}

		for _, event := range events {
			if event.Type == linebot.EventTypeFollow {
				lineId := event.Source.UserID
				query.PostUser(lineId)
			}

			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {

				case *linebot.TextMessage:
					reply.Text(bot, event, message)

				case *linebot.StickerMessage:
					reply.Sticker(bot, event, message)

				case *linebot.ImageMessage:
					reply.Image(bot, event, message)

				}
			}
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
