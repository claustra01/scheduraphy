package query

import (
	"github.com/claustra01/scheduraphy/db"
)

func GetUser(lineId string) *db.User {
	user := new(db.User)
	result := db.Psql.First(&user, "line_id = ?", lineId)
	if result.RowsAffected == 0 {
		return nil
	}
	return user
}
