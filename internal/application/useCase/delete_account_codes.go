package usecase

import (
	"github.com/ropehapi/password-vault-go/internal/domain/entity"
)

type DeleteAccountCodesInputDTO struct {
	Id int64 `json:"id"`
}

type DeleteAccountCodesUseCase struct {
	AccountCodesRepository entity.AccountCodesRepositoryInterface
}

func NewDeleteAccountCodesUseCase(
	AccountCodesRepository entity.AccountCodesRepositoryInterface,
) *DeleteAccountCodesUseCase {
	return &DeleteAccountCodesUseCase{
		AccountCodesRepository: AccountCodesRepository,
	}
}

func (c *DeleteAccountCodesUseCase) Execute(input DeleteAccountCodesInputDTO) error {
	err := c.AccountCodesRepository.Delete(input.Id)
	if err != nil {
		return err
	}
	return nil
}
