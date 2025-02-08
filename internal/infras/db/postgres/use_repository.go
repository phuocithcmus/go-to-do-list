package postgres

import (
	"to_do_project/internal/domain/entities"
	"to_do_project/internal/domain/repositories"

	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) repositories.UserRepository {
	return &GormUserRepository{
		db: db,
	}
}

// Create implements repositories.UserRepository.
func (g *GormUserRepository) Create(user *entities.User) (*entities.User, error) {
	if err := g.db.Create(user).Error; err != nil {
		return nil, err
	}

	// Read row from DB to never return different data than persisted
	return g.FindById(user.Id.String())
}

// FindAll implements repositories.UserRepository.
func (g *GormUserRepository) FindAll() ([]*entities.User, error) {
	var dbUsers []entities.User

	if err := g.db.Preload("Seller").Find(&dbUsers).Error; err != nil {
		return nil, err
	}

	users := make([]*entities.User, len(dbUsers))

	return users, nil
}

// FindById implements repositories.UserRepository.
func (g *GormUserRepository) FindById(id string) (*entities.User, error) {
	var dbUser entities.User
	if err := g.db.Preload("Seller").First(&dbUser, id).Error; err != nil {
		return nil, err
	}

	return &dbUser, nil
}

// FindByUsername implements repositories.UserRepository.
func (g *GormUserRepository) FindByUsername(username string) (*entities.User, error) {
	panic("unimplemented")
}
