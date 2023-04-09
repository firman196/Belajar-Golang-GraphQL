package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"belajar-golang-gql/graph/model"
	"context"
	"fmt"
)

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, input model.InputUser) (*model.TokenOutput, error) {
	panic(fmt.Errorf("not implemented: Register - register"))
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.TokenOutput, error) {
	panic(fmt.Errorf("not implemented: Login - login"))
}

// RefreshToken is the resolver for the refreshToken field.
func (r *mutationResolver) RefreshToken(ctx context.Context) (*model.TokenOutput, error) {
	panic(fmt.Errorf("not implemented: RefreshToken - refreshToken"))
}

// UpdateAccount is the resolver for the updateAccount field.
func (r *mutationResolver) UpdateAccount(ctx context.Context, input *model.UpdateAccountInput) (*model.UserOutput, error) {
	panic(fmt.Errorf("not implemented: UpdateAccount - updateAccount"))
}

// DeleteAccount is the resolver for the deleteAccount field.
func (r *mutationResolver) DeleteAccount(ctx context.Context, id string) (*bool, error) {
	panic(fmt.Errorf("not implemented: DeleteAccount - deleteAccount"))
}

// GetByID is the resolver for the getById field.
func (r *queryResolver) GetByID(ctx context.Context, id string) (*model.UserOutput, error) {
	panic(fmt.Errorf("not implemented: GetByID - getById"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
