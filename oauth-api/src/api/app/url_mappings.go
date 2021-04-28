package app

import (
	"github.com/rjunior/golang-microservice/oauth-api/src/api/controllers/oauth"
	"github.com/rjunior/golang-microservice/src/api/controllers/polo"
)

func mapUrls() {
	router.GET("/marco", polo.Marco)

	router.POST("/oauth/access_token", oauth.CreateAccessToken)
	router.GET("/oauth/access_token/:token_id", oauth.GetAccessToken)
}
