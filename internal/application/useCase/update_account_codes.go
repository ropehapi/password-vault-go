package usecase

import (
	"github.com/ropehapi/password-vault-go/internal/domain/entity"
)

type UpdateAccountCodesUseCase struct {
	AccountCodesRepository entity.AccountCodesRepositoryInterface
}

func NewUpdateAccountCodesUseCase(
	AccountCodesRepository entity.AccountCodesRepositoryInterface,
) *UpdateAccountCodesUseCase {
	return &UpdateAccountCodesUseCase{
		AccountCodesRepository: AccountCodesRepository,
	}
}

func (c *UpdateAccountCodesUseCase) Execute(id int64, input CreateAccountCodesInputDTO) (AccountCodesOutputDTO, error) {
	accountCodes, err := entity.NewAccountCodes(input.Name, input.Codes)
	if err != nil {
		return AccountCodesOutputDTO{}, err
	}

	if err := c.AccountCodesRepository.Update(id, accountCodes); err != nil {
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
