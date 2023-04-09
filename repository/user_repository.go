package repository

import (
	"belajar-golang-gql/graph/model"
	"context"
)

type UserRepository interface {
	//Create(ctx context.Context, user model.InputUser) (*model.UserOutput, error)
	GetById(ctx context.Context, id string) (*model.UserOutput, error)
	/*Update(ctx context.Context, user model.UpdateAccountInput) (model.UserOutput, error)
	Delete(ctx context.Context, id string) bool
	GenerateToken(ctx context.Context, user model.LoginInput) model.TokenOutput
	RefreshToken(ctx context.Context) model.TokenOutput*/
}
