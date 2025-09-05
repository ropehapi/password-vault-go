package usecase

import (
	"github.com/ropehapi/password-vault-go/internal/domain/entity"
	"github.com/ropehapi/password-vault-go/pkg/encrypter"
)

type UpdateAccountUseCase struct {
	AccountRepository entity.AccountRepositoryInterface
}

func NewUpdateAccountUseCase(
	AccountRepository entity.AccountRepositoryInterface,
) *UpdateAccountUseCase {
	return &UpdateAccountUseCase{
		AccountRepository: AccountRepository,
	}
}

func (c *UpdateAccountUseCase) Execute(id int64, input CreateAccountInputDTO) (AccountOutputDTO, error) {
	account, err := c.AccountRepository.GetById(id)
	
	if err != nil {
		return AccountOutputDTO{}, err
	}

	if input.Name != "" {
		account.Name = input.Name
	}

	if input.Login != "" {
		account.Login = input.Login
	}

	if input.Password != "" {
		encrypytedPassword, err := encrypter.Crypt(input.Password)
		if err != nil {
			return AccountOutputDTO{}, err
		}

		account.Password = encrypytedPassword
	}

	if err := c.AccountRepository.Update(id, account); err != nil {
		return AccountOutputDTO{}, err
	}

	dto := AccountOutputDTO{
		ID:        account.ID,
		Name:      account.Name,
		Login:     account.Login,
		Password:  account.Password,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
	}

	return dto, nil
}
