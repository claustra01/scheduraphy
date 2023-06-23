package query

import (
	"errors"
)

func GetRefreshToken(lineId string) (string, error) {
	user := GetUser(lineId)
	if user == nil {
		return "", errors.New("[ERROR] User not found!")
	}
	return user.RefreshToken, nil
}
