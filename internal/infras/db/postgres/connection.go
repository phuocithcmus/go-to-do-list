package postgres

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func NewConnection() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=to_do_project password=postgres sslmode=disable")
	if err != nil {
		return nil, err
	}

	return db, nil
}
