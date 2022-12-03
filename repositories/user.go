package repositories

import (
	"my-gram-1/models/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *entity.User) (*entity.User, error)
	FindByEmail(user entity.User) (entity.User, error)
	Update(user *entity.User) (*entity.User, error)
	Delete(userId uint) error
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserReposittory(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (r *UserRepositoryImpl) Create(user *entity.User) (*entity.User, error) {
	err := r.DB.Create(&user).Error

	if err != nil {
		return nil, err
	}

	return user, err
}

func (r *UserRepositoryImpl) FindByEmail(user entity.User) (entity.User, error) {
	err := r.DB.Model(&user).Where("email = ?", user.Email).First(&user).Error
	return user, err
}

func (r *UserRepositoryImpl) Update(user *entity.User) (*entity.User, error) {
	err := r.DB.Select("username", "email").Updates(&user).First(&user, "id = ?", user.ID).Error
	if err != nil {
		return nil, err
	}

	return user, err
}

func (r *UserRepositoryImpl) Delete(userId uint) error {
	var user entity.User
	err := r.DB.Delete(user, "id = ?", userId).Error

	if err != nil {
		return err
	}

	return err
}
