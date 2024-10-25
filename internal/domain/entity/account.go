package entity

import (
	"errors"
	"github.com/ropehapi/password-vault-go/pkg/encrypter"
	"time"
)

type Account struct {
	ID        string
	Name      string
	Login     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAccount(name, login, password string) (*Account, error) {
	encrypytedPassword, err := encrypter.Crypt(password, []byte("exemplo-chave-32"))
	if err != nil {
		return nil, err
	}

	account := &Account{
		Name:     name,
		Login:    login,
		Password: encrypytedPassword,
	}

	err = account.IsValid()
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (a *Account) IsValid() error {
	if a.Name == "" {
		return errors.New("name is required")
	}
	if a.Login == "" {
		return errors.New("login is required")
	}
	if a.Password == "" {
		return errors.New("password is required")
	}

	return nil
}
