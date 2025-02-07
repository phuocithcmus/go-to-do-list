package interfaces

import "to_do_project/internal/domain/entities"

type UserService interface {
	CreateUser(username string, password string) (*entities.User, error)
	FindAll() ([]*entities.User, error)
}
