package app

import (
	"github.com/gin-gonic/gin"
	"github.com/rjunior/golang-microservice/src/api/log/option_b"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartApp() {
	option_b.Info("Starting mapping", option_b.Field("step", 1), option_b.Field("status", "pending"))
	mapUrls()

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
