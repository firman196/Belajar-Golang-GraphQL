package services

import (
	"belajar-golang-gql/repository"
	"testing"

	"github.com/stretchr/testify/mock"
)

var userRepository = &repository.UserRepositoryMock{Mock: mock.Mock{}}
var userService = UserServiceImpl{UserRepository: userRepository}

func TestUserNotFound(t *testing.T) {

}

/*
var userRepository = &repository.UserRepositoryMock{Mock: mock.Mock{}}
var userService = UserServiceImpl{UserRepository: userRepository}
*/
