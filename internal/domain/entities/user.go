package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID
	Username  string
	Password  string
	CreatedAt time.Time
}

func (u *User) validate() error {
	if u.Username == "" {
		return errors.New("username is empty")
	}

	if u.Password == "" {
		return errors.New("password is empty")
	}

	return nil
}

func NewUser(username string, password string) (*User, error) {
	user := &User{
		Id:        uuid.New(),
		Username:  username,
		Password:  password,
		CreatedAt: time.Now(),
	}

	if err := user.validate(); err != nil {
		return nil, err
	}

	return user, nil
}
