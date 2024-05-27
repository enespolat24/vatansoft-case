package repository

import (
	"vatansoft-case/internal/model"

	"gorm.io/gorm"
)

type roleRepository struct {
	db *gorm.DB
}

type RoleRepository interface {
	GetRole(id string) (*model.Role, error)
	GetRoleByName(name string) (*model.Role, error)
	CreateRole(role *model.Role) error
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{db: db}
}

func (ur *roleRepository) GetRole(id string) (*model.Role, error) {
	var role model.Role
	if err := ur.db.First(&role, id).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (ur *roleRepository) GetRoleByName(name string) (*model.Role, error) {
	var role model.Role
	if err := ur.db.Where("name = ?", name).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (ur *roleRepository) CreateRole(role *model.Role) error {
	return ur.db.Create(role).Error
}
