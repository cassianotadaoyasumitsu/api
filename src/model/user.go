package model

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
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
func (user *User) Preparar(step string) error {
	if err := user.validar(step); err != nil {
		return err
	}

	if err := user.formatar(step); err != nil {
		return err
	}
	return nil
}

func (user *User) validar(step string) error {
	if user.Name == "" {
		return errors.New("name demanded, can't be blank")
	}

	if user.Nick == "" {
		return errors.New("nick demanded, can't be blank")
	}

	if user.Email == "" {
		return errors.New("email demanded, can't be blank")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("email format invalid")
		// return err
	}

	if step == "register" && user.Password == "" {
		return errors.New("password demanded, can't be blank")
	}

	return nil
}

func (user *User) formatar(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if step == "register" {
		hashPassword, err := security.Hash(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(hashPassword)
	}

	return nil
}
