package util

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ayush6624/go-chatgpt"
)

func FormatJson() {

	key := os.Getenv("OPENAI_KEY")
	c, err := chatgpt.NewClient(key)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	res, err := c.SimpleSend(ctx, "Hey, Explain GoLang to me in 2 sentences.")
	if err != nil {
		log.Fatal(err)
	}

	msg := res.Choices[0].Message.Content
	fmt.Print(msg)

}
