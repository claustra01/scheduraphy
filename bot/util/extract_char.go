package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func ExtractChar(imageData []byte) string {

	key := os.Getenv("AZURE_COMPUTER_VISION_KEY")
	endpoint := os.Getenv("AZURE_COMPUTER_VISION_ENDPOINT")

	// 画像データをリクエストボディに設定
	body := bytes.NewReader(imageData)

	// REST APIの実行
	uri := endpoint + "/vision/v2.1/ocr?language=unk&detectOrientation=true"
	req, _ := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Add("Ocp-Apim-Subscription-Key", key)
	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	// JSONに展開
	var data interface{}
	dataStr, _ := ioutil.ReadAll(resp.Body)
	err := json.Unmarshal([]byte(dataStr), &data)
	if err != nil {
		fmt.Println("JSON parsing error:", err)
		return ""
	}

	// 文字列に整形
	imageStr := ""
	regions := data.(map[string]interface{})["regions"].([]interface{})
	for _, region := range regions {
		lines := region.(map[string]interface{})["lines"].([]interface{})
		for _, line := range lines {
			words := line.(map[string]interface{})["words"].([]interface{})
			for _, word := range words {
				text := word.(map[string]interface{})["text"].(string)
				imageStr += text
			}
			imageStr += "\n"
		}
	}

	return imageStr

}
