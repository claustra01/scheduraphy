package token

import (
	"github.com/claustra01/scheduraphy/db"
)

func GetRefreshToken(lineId string) string {
	user := new(db.User)
	db.Psql.First(&user, "line_id = ?", lineId)
	return user.RefreshToken
}
