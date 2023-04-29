package services

import (
	"belajar-golang-gql/graph/model"
	"belajar-golang-gql/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository = &repository.UserRepositoryMock{Mock: mock.Mock{}}
var userService = UserServiceImpl{UserRepository: userRepository}

type UserOutput struct {
	UserId    string
	Firstname string
	Lastname  *string
	Email     string
}

// Negative case
func TestUserNotFound(t *testing.T) {
	userRepository.Mock.On("GetById", "1").Return(nil)
	user, err := userService.GetById("1")
	assert.Nil(t, user)
	assert.NotNil(t, err)
}

// Positive case
func TestUserSucess(t *testing.T) {
	user := model.UserOutput{
		UserID:    "2",
		Firstname: "Firman Agus",
		Lastname:  "Tes",
		Email:     "nnheo@example.com",
	}
	userRepository.Mock.On("GetById", "2").Return(user)
	response, err := userService.GetById("2")
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, user.Firstname, response.Firstname)
	assert.Equal(t, user.Lastname, response.Lastname)
	assert.Equal(t, user.Email, response.Email)
}
