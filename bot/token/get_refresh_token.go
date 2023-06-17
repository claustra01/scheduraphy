package token

import (
	"log"

	"github.com/claustra01/scheduraphy/db"
)

func GetRefreshToken(lineId string) string {
	user := new(db.User)
	result := db.Psql.First(&user, "line_id = ?", lineId)
	if result.RowsAffected == 0 {
		log.Print("[INFO] User not found!")
		return ""
	}
	return user.RefreshToken
}
