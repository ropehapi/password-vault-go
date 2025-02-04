package usecase

import (
	"github.com/ropehapi/password-vault-go/internal/domain/entity"
	"time"
)

type CreateAccountCodesInputDTO struct {
	Name  string `json:"name"`
	Codes string `json:"codes"`
}

type AccountCodesOutputDTO struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Codes     string    `json:"codes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateAccountCodesUseCase struct {
	AccountCodesRepository entity.AccountCodesRepositoryInterface
}

func NewCreateAccountCodesUseCase(
	AccountCodesRepository entity.AccountCodesRepositoryInterface,
) *CreateAccountCodesUseCase {
	return &CreateAccountCodesUseCase{
		AccountCodesRepository: AccountCodesRepository,
	}
}

func (c *CreateAccountCodesUseCase) Execute(input CreateAccountCodesInputDTO) (AccountCodesOutputDTO, error) {
	accountCodes, err := entity.NewAccountCodes(input.Name, input.Codes)
	if err != nil {
		return AccountCodesOutputDTO{}, err
	}

	if err := c.AccountCodesRepository.Save(accountCodes); err != nil {
		return AccountCodesOutputDTO{}, err
	}

	dto := AccountCodesOutputDTO{
		ID:        accountCodes.ID,
		Name:      accountCodes.Name,
		Codes:     accountCodes.Codes,
		CreatedAt: accountCodes.CreatedAt,
		UpdatedAt: accountCodes.UpdatedAt,
	}

	return dto, nil
}
