package repository

import (
	"belajar-golang-gql/graph/model"
	"errors"

	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (repository *UserRepositoryMock) GetById(ctx context.Context, id string) (*model.UserOutput, error) {
	mockTest := repository.Mock.Called(id)

	if mockTest.Get(0) != nil {
		user := mockTest.Get(0).(model.UserOutput)
		return &user, nil
	} else {
		return nil, errors.New("internal server error ")
	}
}
