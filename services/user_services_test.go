package services

import (
	"belajar-golang-gql/repository"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository = &repository.UserRepositoryMock{Mock: mock.Mock{}}
var userService = UserServiceImpl{UserRepository: userRepository}

func TestUserNotFound(t *testing.T) {
	userRepository.Mock.On("GetById", "1").Return(nil)
	ctx := context.Background()
	user, err := userService.GetById(ctx, "1")
	assert.Nil(t, user)
	assert.NotNil(t, err)
}
