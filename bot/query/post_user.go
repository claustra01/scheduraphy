package query

import (
	"log"
	"time"

	"github.com/claustra01/scheduraphy/db"
	"github.com/google/uuid"
)

func PostUser(lineId string) *db.User {

	if user := GetUser(lineId); user != nil {
		log.Print("[INFO] User already exists!")
		return user
	}

	user := new(db.User)
	user.Id, _ = uuid.NewUUID()
	user.LineId = lineId
	user.RefreshToken = ""
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	db.Psql.Create(&user)
	return user
}
