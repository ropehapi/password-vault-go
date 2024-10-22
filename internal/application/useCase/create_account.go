package usecase

import (
	"github.com/ropehapi/password-vault-go/internal/domain/entity"
	"time"
)

type CreateAccountInputDTO struct {
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type AccountOutputDTO struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateAccountUseCase struct {
	AccountRepository entity.AccountRepositoryInterface
}

func NewCreateAccountUseCase(
	AccountRepository entity.AccountRepositoryInterface,
) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		AccountRepository: AccountRepository,
	}
}

func (c *CreateAccountUseCase) Execute(input CreateAccountInputDTO) (AccountOutputDTO, error) {
	account, err := entity.NewAccount(input.Name, input.Login, input.Password)
	if err != nil {
		return AccountOutputDTO{}, err
	}

	if err := c.AccountRepository.Save(account); err != nil {
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
