package services

import (
	"to_do_project/internal/application/interfaces"
	"to_do_project/internal/domain/entities"
	"to_do_project/internal/domain/repositories"
)

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) interfaces.UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) CreateUser(username string, password string) (*entities.User, error) {
	user, err := s.userRepository.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	if user != nil {
		return nil, err
	}

	var newUser, errNewUser = entities.NewUser(username, password)

	if errNewUser != nil {
		return nil, errNewUser
	}

	user, err = s.userRepository.Create(newUser)
	if err != nil {
		return nil, err
	}

	return user, nil

}

func (s *UserService) FindAll() ([]*entities.User, error) {
	user, err := s.userRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return user, nil
}
