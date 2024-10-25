package usecase

import (
	"github.com/ropehapi/password-vault-go/internal/domain/entity"
)

type DeleteAccountInputDTO struct {
	Id int64 `json:"id"`
}

type DeleteAccountUseCase struct {
	AccountRepository entity.AccountRepositoryInterface
}

func NewDeleteAccountUseCase(
	AccountRepository entity.AccountRepositoryInterface,
) *DeleteAccountUseCase {
	return &DeleteAccountUseCase{
		AccountRepository: AccountRepository,
	}
}

func (c *DeleteAccountUseCase) Execute(input DeleteAccountInputDTO) error {
	err := c.AccountRepository.Delete(input.Id)
	if err != nil {
		return err
	}
	return nil
}
