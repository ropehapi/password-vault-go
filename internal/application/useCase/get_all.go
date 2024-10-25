package usecase

import (
	"github.com/ropehapi/password-vault-go/internal/domain/entity"
	"github.com/ropehapi/password-vault-go/pkg/encrypter"
)

type GetAllAccountsUseCase struct {
	AccountRepository entity.AccountRepositoryInterface
}

func NewGetAllAccountsUseCase(
	AccountRepository entity.AccountRepositoryInterface,
) *GetAllAccountsUseCase {
	return &GetAllAccountsUseCase{
		AccountRepository: AccountRepository,
	}
}

func (c *GetAllAccountsUseCase) Execute() ([]AccountOutputDTO, error) {
	accounts, err := c.AccountRepository.GetAll()
	if err != nil {
		return nil, err
	}

	var accountDTOs []AccountOutputDTO

	for _, account := range accounts {
		decryptedString, err := encrypter.Decrypt(account.Password, []byte("exemplo-chave-32"))
		if err != nil {
			return nil, err
		}

		dto := AccountOutputDTO{
			ID:        account.ID,
			Name:      account.Name,
			Login:     account.Login,
			Password:  decryptedString,
			CreatedAt: account.CreatedAt,
			UpdatedAt: account.UpdatedAt,
		}

		accountDTOs = append(accountDTOs, dto)
	}

	return accountDTOs, nil
}
