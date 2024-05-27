package repository

import (
	"vatansoft-case/internal/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *model.User) error
	GetUsers() ([]model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) CreateUser(user *model.User) error {
	return ur.db.Create(user).Error
}

func (ur *userRepository) GetUsers() ([]model.User, error) {
	var users []model.User
	if err := ur.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
