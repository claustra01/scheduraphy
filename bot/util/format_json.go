package util

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/ayush6624/go-chatgpt"
)

func FormatJson(rawStr string) interface{} {

	key := os.Getenv("OPENAI_KEY")
	c, err := chatgpt.NewClient(key)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	str := fmt.Sprintf(`以下に示す文章はイベントのポスターの写真や告知のスクリーンショットなどから画像認識AIを用いて文字のみを抽出したものです．この文章から明らかに意味がないものを除去してください．また，画像認識AIの精度が良くないため，文法的あるいは語彙的に不明瞭なものがあれば適宜修正してください．イベント名，日付と時刻，会場と住所の内容に関しては分かりやすいように明記し，不必要な改行は除去してください．
ただし，修正内容についての説明などは不要です．修正後の文章以外の内容を出力した場合，あなたの過失によって無関係の人々に危害を加えてしまう可能性があります．
----------
%s`, rawStr)

	res, err := c.SimpleSend(ctx, str)
	if err != nil {
		log.Fatal(err)
	}

	str = fmt.Sprintf(`イベントのポスターの写真や告知のスクリーンショットなどから画像認識AIを用いて文字のみを出力した文章を渡します．この文章を以下に示す例に則ったJSON形式に整形してください．ただし，\"venue\"に対応するものが文章中に見つからなかった場合，そのキーを省略したりせず，値を空文字列にしてください．\n----------
{"type":"event","title":"福岡県大会","start":"2023-06-20T10:00:00+09:00","end":"2023-06-21T18:00:00+09:00","location":"福岡県立体育館"}
----------
渡された文章がイベントに関するものでないと判断した場合は {"type":"null"} のみを出力してください．ただし，最初の数行はイベントに関係するものではない無関係な文字列である可能性を考慮した上で，最初の行を無視して再帰的に判断してください．
また，JSONの内容についての説明などは不要です．JSON以外の内容を出力した場合，あなたの過失によって無関係の人々に危害を加えてしまう可能性があります．
----------
%s`, res.Choices[0].Message.Content)

	res, err = c.SimpleSend(ctx, str)
	if err != nil {
		log.Fatal(err)
	}

	var data interface{}
	err = json.Unmarshal([]byte(res.Choices[0].Message.Content), &data)
	if err != nil {
		log.Print(err)
	}

	return data
}
