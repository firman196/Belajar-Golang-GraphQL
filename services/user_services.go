package services

import (
	"belajar-golang-gql/graph/model"
	"context"
)

type UserServices interface {
	//Register(ctx context.Context, input model.InputUser) (model.TokenOutput, error)
	GetById(ctx context.Context, userId string) (*model.UserOutput, error)
}
