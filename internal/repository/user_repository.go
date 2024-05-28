package repository

import (
	"vatansoft-case/internal/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *model.User) error
	GetUsers() ([]model.User, error)
	UserExistsByEmail(email string) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	GetUserById(id uint) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id string) error
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

func (ur *userRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := ur.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepository) UserExistsByEmail(email string) (*model.User, error) {
	var user model.User
	if err := ur.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepository) GetUserById(id uint) (*model.User, error) {
	var user model.User
	if err := ur.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepository) DeleteUser(id string) error {
	user := &model.User{}
	result := ur.db.First(user, id)
	if result.Error != nil {
		return result.Error
	}

	if err := ur.db.Delete(user).Error; err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) UpdateUser(user *model.User) error {
	return ur.db.Save(user).Error
}
