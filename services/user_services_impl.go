package services

import (
	"belajar-golang-gql/graph/model"
	"belajar-golang-gql/helper"
	"belajar-golang-gql/models/entity"
	"belajar-golang-gql/models/web"
	"belajar-golang-gql/repository"
	"belajar-golang-gql/utils"
	"errors"
	"os"
	"strconv"
	"time"

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

func (service *UserServiceImpl) GetById(id string) (*entity.UserRepository, error) {
	user, err := service.UserRepository.GetById(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil

}

func (service *UserServiceImpl) Create(input model.InputUser) (*model.TokenOutput, error) {
	passwordHash, err := helper.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	user := entity.UserRepository{
		UserID:    utils.Uuid(),
		Firstname: input.Firstname,
		Lastname:  input.Lastname,
		Email:     input.Email,
		Password:  passwordHash,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	response, err := service.UserRepository.Create(user)
	if err != nil {
		return nil, err
	}

	jwtExpiredTimeToken, _ := strconv.Atoi(os.Getenv("JWT_EXPIRED_TIME_TOKEN"))
	jwtExpiredTimeRefreshToken, _ := strconv.Atoi(os.Getenv("JWT_EXPIRED_TIME_REFRESH_TOKEN"))
	tokenCreateRequest := web.TokenCreateRequest{
		UserID:    response.UserID,
		Firstname: response.Firstname,
		Lastname:  response.Lastname,
		Email:     response.Email,
	}

token:
	web.TokenResponse{
		Token: helper.CreateToken(tokenCreateRequest),
	}
}
