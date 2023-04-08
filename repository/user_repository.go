package repository

import (
	"belajar-golang-gql/graph/model"
	"context"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, db *gorm.DB, user model.InputUser) (*model.UserOutput, error)
	FindByID(ctx context.Context, db *gorm.DB, id string) model.UserOutput
	Update(ctx context.Context, db *gorm.DB, user model.UpdateAccountInput) model.UserOutput
	Delete(ctx context.Context, db *gorm.DB, id string) bool
	GenerateToken(ctx context.Context, db *gorm.DB, user model.LoginInput) model.TokenOutput
	RefreshToken(ctx context.Context, db *gorm.DB) model.TokenOutput
}
