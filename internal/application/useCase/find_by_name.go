package usecase

import (
	"github.com/ropehapi/password-vault-go/internal/domain/entity"
)

type FindAccountByNameInputDTO struct {
	Name string `json:"name"`
}

type FindAccountByNameUseCase struct {
	AccountRepository entity.AccountRepositoryInterface
}

func NewFindAccountByNameUseCase(
	AccountRepository entity.AccountRepositoryInterface,
) *FindAccountByNameUseCase {
	return &FindAccountByNameUseCase{
		AccountRepository: AccountRepository,
	}
}

func (c *FindAccountByNameUseCase) Execute(input FindAccountByNameInputDTO) (AccountOutputDTO, error) {
	account, err := c.AccountRepository.FindByName(input.Name)
	if err != nil {
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
