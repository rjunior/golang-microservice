package oauth

import (
	"github.com/rjunior/golang-microservice/src/api/utils/errors"
)

const (
	QueryGetUserByUsernameAndPassword = "SELECT id, username FROM users WHERE username=? AND passrod=?;"
)

var (
	users = map[string]*User{
		"fede": {Id: 123, Username: "fede"},
	}
)

func GetUserByUsernameAndPassword(username, password string) (*User, errors.ApiError) {
	user := users[username]
	if user == nil {
		return nil, errors.NewNotFoundError("no user found with given parameters")
	}

	return user, nil
}
