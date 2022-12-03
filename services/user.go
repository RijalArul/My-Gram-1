package services

import (
	"my-gram-1/models/entity"
	"my-gram-1/repositories"
)

type UserService interface {
	Register(userInput *entity.User) (*entity.User, error)
	Login(userInput entity.User) (entity.User, error)
	Profile(userInput entity.User) (entity.User, error)
	Update(userInput *entity.User) (*entity.User, error)
	Delete(userId uint) error
}

type UserServiceImpl struct {
	UserRepository repositories.UserRepository
	UserEntity     entity.User
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepository: userRepo,
	}
}

func (s *UserServiceImpl) Register(userInput *entity.User) (*entity.User, error) {
	newUser, err := s.UserRepository.Create(userInput)
	return newUser, err
}

func (s *UserServiceImpl) Login(userInput entity.User) (entity.User, error) {
	loginUser, err := s.UserRepository.FindByEmail(userInput)
	return loginUser, err
}

func (s *UserServiceImpl) Profile(userInput entity.User) (entity.User, error) {
	getUser, err := s.UserRepository.FindByEmail(userInput)
	return getUser, err
}

func (s *UserServiceImpl) Update(userInput *entity.User) (*entity.User, error) {
	updateUser, err := s.UserRepository.Update(userInput)
	return updateUser, err
}

func (s *UserServiceImpl) Delete(userId uint) error {
	err := s.UserRepository.Delete(userId)
	return err
}
