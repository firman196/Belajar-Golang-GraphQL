package repository

import (
	"belajar-golang-gql/models/entity"
	"belajar-golang-gql/utils"
)

type UserRepository interface {
	Create(user entity.UserRepository) (*entity.UserRepository, error)
	GetById(id string) (*entity.UserRepository, error)
	Update(cuser entity.UserRepository) (*entity.UserRepository, error)
	Delete(id string) (bool, error)
	GetList(pagination utils.Pagination) (*utils.Pagination, error)
}
