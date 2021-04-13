package domain

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rjunior/golang-microservice/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Fede", Lastname: "Leon", Email: "myemail@gmail.com"},
	}

	UserDao usersDaoInterface
)

func init() {
	UserDao = &userDao{}
}

type usersDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct{}

func (u *userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {
	log.Println("we're accessing the database")
	if user := users[userId]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
