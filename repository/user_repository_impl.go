package repository

import (
	"belajar-golang-gql/graph/model"
	"context"
	"errors"

	"belajar-golang-gql/graph/utils"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct{}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func Create(ctx context.Context, db *gorm.DB, user model.InputUser) (*model.UserOutput, error) {
	if err := db.Save(&user).Error; err != nil {
		utils.SetFatalError(err)
		return nil, errors.New("Internal Server Error ")
	}

}
