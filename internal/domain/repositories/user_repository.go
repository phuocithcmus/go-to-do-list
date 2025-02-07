package repositories

import "to_do_project/internal/domain/entities"

type UserRepository interface {
	Create(user *entities.User) (*entities.User, error)
	FindByUsername(username string) (*entities.User, error)
	FindById(id string) (*entities.User, error)
	FindAll() ([]*entities.User, error)
}
