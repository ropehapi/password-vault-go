package controllers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	usecase "github.com/ropehapi/password-vault-go/internal/application/useCase"
	"github.com/ropehapi/password-vault-go/internal/domain/entity"

	"net/http"
)

type CreateAccountController struct {
	AccountRepository entity.AccountRepositoryInterface
}

func NewCreateAccountController(accountRepository entity.AccountRepositoryInterface) *CreateAccountController {
	return &CreateAccountController{
		AccountRepository: accountRepository,
	}
}

func (c *CreateAccountController) Create(w http.ResponseWriter, r *http.Request) {
	var dto usecase.CreateAccountInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createAccountUseCase := usecase.NewCreateAccountUseCase(c.AccountRepository)
	output, err := createAccountUseCase.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
	return
}

func (c *CreateAccountController) FindByName(w http.ResponseWriter, r *http.Request) {
	dto := usecase.FindAccountByNameInputDTO{
		Name: chi.URLParam(r, "name"),
	}

	findAccountByNameUseCase := usecase.NewFindAccountByNameUseCase(c.AccountRepository)
	output, err := findAccountByNameUseCase.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
	return
}
