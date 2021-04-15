package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserfound(t *testing.T) {
	// Initialization:

	// Execution:
	user, err := UserDao.GetUser(0)

	// Validation:
	assert.Nil(t, user, "We were not expeting a user with id 0")
	assert.NotNil(t, err, "We were expecting a error when user id is 0")
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "user 0 was not found", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	user, err := UserDao.GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)

	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "Fede", user.FirstName)
	assert.EqualValues(t, "Leon", user.Lastname)
	assert.EqualValues(t, "myemail@gmail.com", user.Email)
}
