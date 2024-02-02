package db

import (
	"errors"
	"go-jwt/db/models"
	"golang-jwt/db/models"
)

var users = map[string]models.User{}

func StoreUser(userName, Password, role string) (string, error) {}

func InitDB() {

}

func FetchUserByUsername(userName string) (models.User, string, error) {

	for k, v := range users {
		if v.UserName == userName {
			return v, k, nil
		}
	}

	return models.User{}, "", errors.New("user not found that mathches the given username")

}

func DeleteUser() {}

func FetchUserByID() {}

func StoreRefreshToken()

func DeleteRefreshToken()

func CheckRefreshToken() bool {

}

func LogUserIn() {}

func genrateBcryptHash() {}

func CheckPasswordAgainstHash() error {

}
