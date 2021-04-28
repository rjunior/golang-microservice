package oauth

import (
	"fmt"

	"github.com/rjunior/golang-microservice/src/api/utils/errors"
)

var (
	tokens = make(map[string]*AccessToken, 0)
)

func (at *AccessToken) Save() errors.ApiError {
	at.AccessToken = fmt.Sprintf("USR_%d", at.UserId) // Create a random string to me the access token
	tokens[at.AccessToken] = at
	return nil
}

func GetAccessTokenByToken(accessToken string) (*AccessToken, errors.ApiError) {
	token := tokens[accessToken]
	if token == nil {
		return nil, errors.NewNotFoundError("no access token with given parameters")
	}

	return token, nil
}
