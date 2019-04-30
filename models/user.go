package models

import (
	"errors"

	"github.com/go-redis/redis"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserNotFound = errors.New("User Not Found")
	ErrInvalidAuth  = errors.New("Invalid Authentication")
)

type user struct {
	key string
}

func newUser(username string, hash []byte) (*user, error) {
	id, err := db.Incr("user:next-id").Result()
	if err != nil {
		return nil, err
	}

}

func Login(username, password string) error {
	hash, err := db.Get("user:" + username).Bytes()

	if err == redis.Nil {
		return ErrUserNotFound
	} else if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword(hash, []byte(password))

	if err != nil {
		return ErrInvalidAuth
	}

	return nil
}

func Register(username, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	return db.Set("user:"+username, hash, 0).Err()
}
