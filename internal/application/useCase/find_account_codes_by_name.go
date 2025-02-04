package usecase

import (
	"github.com/ropehapi/password-vault-go/internal/domain/entity"
	"github.com/ropehapi/password-vault-go/pkg/encrypter"
)

type GetAccountCodesByNameInputDTO struct {
	Name string `json:"name"`
}

type GetAccountCodesByNameUseCase struct {
	AccountCodesRepository entity.AccountCodesRepositoryInterface
}

func NewGetAccountCodesByNameUseCase(
	AccountCodesRepository entity.AccountCodesRepositoryInterface,
) *GetAccountCodesByNameUseCase {
	return &GetAccountCodesByNameUseCase{
		AccountCodesRepository: AccountCodesRepository,
	}
}

func (c *GetAccountCodesByNameUseCase) Execute(input GetAccountCodesByNameInputDTO) ([]AccountCodesOutputDTO, error) {
	accounts, err := c.AccountCodesRepository.GetByName(input.Name)
	if err != nil {
		return nil, err
	}

	var accountCodesDTOs []AccountCodesOutputDTO

	for _, account := range accounts {
		decryptedString, err := encrypter.Decrypt(account.Codes)
		if err != nil {
			return nil, err
		}

		dto := AccountCodesOutputDTO{
			ID:        account.ID,
			Name:      account.Name,
			Codes:     decryptedString,
			CreatedAt: account.CreatedAt,
			UpdatedAt: account.UpdatedAt,
		}

		accountCodesDTOs = append(accountCodesDTOs, dto)
	}

	return accountCodesDTOs, nil
}
