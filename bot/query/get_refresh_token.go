package query

import (
	"log"
)

func GetRefreshToken(lineId string) string {
	user := GetUser(lineId)
	if user == nil {
		log.Print("[INFO] User not found!")
		return ""
	}
	return user.RefreshToken
}
