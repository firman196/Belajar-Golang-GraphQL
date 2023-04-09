package repository

import (
	"belajar-golang-gql/graph/model"
	"context"
	"errors"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepositoryImpl(DB *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		DB: DB,
	}
}

/*
func Create(ctx context.Context, db *gorm.DB, user model.InputUser) (*model.UserOutput, error) {
	if err := db.Save(&user).Error; err != nil {
		return nil, errors.New("internal server error ")
	}
	userOutput := &model.UserOutput{
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
	}
	return userOutput, nil
}*/

func (repository *UserRepositoryImpl) GetById(ctx context.Context, id string) (*model.UserOutput, error) {
	user := model.InputUser{}
	if err := repository.DB.First(&user).Error; err != nil {
		return nil, errors.New("internal server error ")
	}
	userOutput := &model.UserOutput{
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
	}
	return userOutput, nil
}
