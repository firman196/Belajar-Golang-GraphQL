package services

import (
	"belajar-golang-gql/graph/model"
	"belajar-golang-gql/repository"
	"context"
	"errors"

	"gorm.io/gorm"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserServiceImpl(userRepository repository.UserRepository, db *gorm.DB) UserServices {
	return &UserServiceImpl{
		UserRepository: userRepository,
	}
}

func (service *UserServiceImpl) GetById(ctx context.Context, id string) (*model.UserOutput, error) {
	user, err := service.UserRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil

}

/*
func (service *UserServiceImpl) Register(ctx context.Context, input model.InputUser) (model.TokenOutput, error) {

}*/
