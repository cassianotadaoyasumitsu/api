package model

import (
	"errors"
	"strings"
	"time"
)

// User model
type User struct {
	ID         uint64    `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	Nick       string    `json:"nick,omitempty"`
	Email      string    `json:"email,omitempty"`
	Password   string    `json:"password,omitempty"`
	Created_at time.Time `json:"created_at,omitempty"`
}

// Chama metodods de validação do Usuário
func (user *User) Preparar() error {
	if err := user.validar(); err != nil {
		return err
	}

	user.formatar()
	return nil
}

func (user *User) validar() error {
	if user.Name == "" {
		return errors.New("name demanded, can't be blank")
	}

	if user.Nick == "" {
		return errors.New("nick demanded, can't be blank")
	}

	if user.Email == "" {
		return errors.New("email demanded, can't be blank")
	}

	if user.Password == "" {
		return errors.New("password demanded, can't be blank")
	}

	return nil
}

func (user *User) formatar() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}
