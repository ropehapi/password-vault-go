package usecase

import (
	"github.com/ropehapi/password-vault-go/internal/domain/entity"
	"github.com/ropehapi/password-vault-go/pkg/encrypter"
)

type GetAccountByNameInputDTO struct {
	Name string `json:"name"`
}

type GetAccountByNameUseCase struct {
	AccountRepository entity.AccountRepositoryInterface
}

func NewGetAccountByNameUseCase(
	AccountRepository entity.AccountRepositoryInterface,
) *GetAccountByNameUseCase {
	return &GetAccountByNameUseCase{
		AccountRepository: AccountRepository,
	}
}

func (c *GetAccountByNameUseCase) Execute(input GetAccountByNameInputDTO) ([]AccountOutputDTO, error) {
	accounts, err := c.AccountRepository.GetByName(input.Name)
	if err != nil {
		return nil, err
	}

	var accountDTOs []AccountOutputDTO

	for _, account := range accounts {
		decryptedString, err := encrypter.Decrypt(account.Password)
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

		// Adiciona o DTO ao slice de resultados
		accountDTOs = append(accountDTOs, dto)
	}

	return accountDTOs, nil
}
