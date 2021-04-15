package app

import (
	"github.com/rjunior/golang-microservice/mvc/controllers"
)

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}
