package repository

import (
	"belajar-golang-gql/models/entity"
	"belajar-golang-gql/utils"
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

// create new data user
// parameter ctx context, db gorm , user structure
// return user structure & error response
func (repository *UserRepositoryImpl) Create(user entity.UserRepository) (*entity.UserRepository, error) {
	if err := repository.DB.Save(&user).Error; err != nil {
		return nil, errors.New("internal server error ")
	}
	return &user, nil
}

// Get user data by user_id
func (repository *UserRepositoryImpl) GetById(id string) (*entity.UserRepository, error) {
	user := entity.UserRepository{}
	if err := repository.DB.First(&user, "user_id = ?", id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, errors.New("user with id = " + id + " not found")
		default:
			return nil, errors.New("internal server error ")
		}
	}

	return &user, nil
}

// update data user
func (repository *UserRepositoryImpl) Update(user entity.UserRepository) (*entity.UserRepository, error) {
	if err := repository.DB.Model(&user.Firstname).Where("user_id = ?", user.UserID).Updates(map[string]interface{}{"user_id": user.UserID, "firstname": user.Firstname, "lastname": user.Lastname, "email": user.Email}).Error; err != nil {
		return nil, errors.New("internal server error ")
	}
	return &user, nil
}

// delete data user
func (repository *UserRepositoryImpl) Delete(id string) (bool, error) {
	var user entity.UserRepository
	if err := repository.DB.Where("user_id = ?", id).Delete(&user).Error; err != nil {
		return false, err
	}
	return true, nil
}

// Get List User
func (repository *UserRepositoryImpl) GetList(pagination utils.Pagination) (*utils.Pagination, error) {
	var users []entity.UserRepository
	repository.DB.Scopes(pagination.Paginate(users, repository.DB)).Find(&users)
	pagination.Rows = users
	return &pagination, nil
}
