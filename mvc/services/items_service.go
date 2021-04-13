package services

import (
	"net/http"

	"github.com/rjunior/golang-microservice/mvc/domain"
	"github.com/rjunior/golang-microservice/mvc/utils"
)

type itemsService struct {
}

var (
	ItemsService itemsService
)

func GetItem(itemId string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "implement me",
		StatusCode: http.StatusInternalServerError,
	}
}
