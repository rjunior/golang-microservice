package services

import (
	"github.com/rjunior/golang-microservice/mvc/domain"
	"github.com/rjunior/golang-microservice/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
