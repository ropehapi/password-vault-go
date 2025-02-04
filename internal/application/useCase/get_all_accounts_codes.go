package usecase

import (
	"github.com/ropehapi/password-vault-go/internal/domain/entity"
	"github.com/ropehapi/password-vault-go/pkg/encrypter"
)

type GetAllAccountsCodesUseCase struct {
	AccountCodesRepository entity.AccountCodesRepositoryInterface
}

func NewGetAllAccountsCodesUseCase(
	AccountCodesRepository entity.AccountCodesRepositoryInterface,
) *GetAllAccountsCodesUseCase {
	return &GetAllAccountsCodesUseCase{
		AccountCodesRepository: AccountCodesRepository,
	}
}

func (c *GetAllAccountsCodesUseCase) Execute() ([]AccountCodesOutputDTO, error) {
	accountsCodes, err := c.AccountCodesRepository.GetAll()
	if err != nil {
		return nil, err
	}

	var accountDTOs []AccountCodesOutputDTO

	for _, account := range accountsCodes {
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

		accountDTOs = append(accountDTOs, dto)
	}

	return accountDTOs, nil
}
