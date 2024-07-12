package service

import (
	"github.com/Anu-renjith/gin-sqlx/entity"
	"github.com/Anu-renjith/gin-sqlx/repository"
)

type UserService interface {
	CreateUser(user entity.User) (entity.User, error)
	GetAllUsers() ([]entity.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepository: userRepo}
}

func (s *userService) CreateUser(user entity.User) (entity.User, error) {
	return s.userRepository.Save(user)
}

func (s *userService) GetAllUsers() ([]entity.User, error) {
	return s.userRepository.FindAll()
}
