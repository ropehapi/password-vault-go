package controllers

import (
	"encoding/json"
	usecase "github.com/ropehapi/password-vault-go/internal/application/useCase"
	"github.com/ropehapi/password-vault-go/internal/domain/entity"
	"net/http"
)

type AccountCodesController struct {
	AccountCodesRepository entity.AccountCodesRepositoryInterface
}

func NewAccountCodesController(accountCodesRepository entity.AccountCodesRepositoryInterface) *AccountCodesController {
	return &AccountCodesController{
		AccountCodesRepository: accountCodesRepository,
	}
}

func (c *AccountCodesController) Create(w http.ResponseWriter, r *http.Request) {
	var dto usecase.CreateAccountCodesInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		response := APIResponse{
			Message: "Erro ao criar conta: " + err.Error(),
			Data:    nil,
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	createAccountCodesUseCase := usecase.NewCreateAccountCodesUseCase(c.AccountCodesRepository)
	output, err := createAccountCodesUseCase.Execute(dto)
	if err != nil {
		response := APIResponse{
			Message: "Erro ao criar conta: " + err.Error(),
			Data:    nil,
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	response := APIResponse{
		Message: "Conta criada com sucesso",
		Data:    output,
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
