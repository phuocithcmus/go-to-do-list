package postgres

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `gorm:primaryKey`
	Username  string
	Password  string
	CreatedAt time.Time
}
