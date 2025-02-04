package entity

import (
	"errors"
	"github.com/ropehapi/password-vault-go/pkg/encrypter"
	"time"
)

type AccountCodes struct {
	ID        string
	Name      string
	Codes     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAccountCodes(name, codes string) (*AccountCodes, error) {
	encrypytedCodes, err := encrypter.Crypt(codes)
	if err != nil {
		return nil, err
	}

	accountCodes := &AccountCodes{
		Name:  name,
		Codes: encrypytedCodes,
	}

	err = accountCodes.IsValid()
	if err != nil {
		return nil, err
	}

	return accountCodes, nil
}

func (a *AccountCodes) IsValid() error {
	if a.Name == "" {
		return errors.New("name is required")
	}
	if a.Codes == "" {
		return errors.New("codes is required")
	}

	return nil
}
