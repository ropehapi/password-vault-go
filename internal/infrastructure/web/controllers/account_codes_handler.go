package controllers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	usecase "github.com/ropehapi/password-vault-go/internal/application/useCase"
	"github.com/ropehapi/password-vault-go/internal/domain/entity"
	"net/http"
	"strconv"
)

type AccountCodesController struct {
	AccountCodesRepository entity.AccountCodesRepositoryInterface
}

func NewAccountCodesController(accountCodesRepository entity.AccountCodesRepositoryInterface) *AccountCodesController {
	return &AccountCodesController{
		AccountCodesRepository: accountCodesRepository,
	}
}

func (c *AccountCodesController) GetAll(w http.ResponseWriter, r *http.Request) {
	getAllAccountsUsecase := usecase.NewGetAllAccountsCodesUseCase(c.AccountCodesRepository)
	output, err := getAllAccountsUsecase.Execute()
	if err != nil {
		response := APIResponse{
			Message: "Erro ao buscar contas: " + err.Error(),
			Data:    nil,
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	if len(output) == 0 {
		response := APIResponse{
			Message: "Não foram encontradas contas",
			Data:    nil,
		}
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	response := APIResponse{
		Message: "Contas encontradas com sucesso",
		Data:    output,
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (c *AccountCodesController) GetByName(w http.ResponseWriter, r *http.Request) {
	dto := usecase.GetAccountCodesByNameInputDTO{
		Name: chi.URLParam(r, "name"),
	}

	findAccountByNameUseCase := usecase.NewGetAccountCodesByNameUseCase(c.AccountCodesRepository)
	output, err := findAccountByNameUseCase.Execute(dto)
	if err != nil {
		response := APIResponse{
			Message: "Erro ao buscar a conta: " + err.Error(),
			Data:    nil,
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	if len(output) == 0 {
		response := APIResponse{
			Message: "Nenhuma conta encontrada para o nome fornecido",
			Data:    nil,
		}
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	response := APIResponse{
		Message: "Contas encontradas com sucesso",
		Data:    output,
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
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

func (c *AccountCodesController) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	idInt, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response := APIResponse{
			Message: "Erro ao deletar conta: " + err.Error(),
			Data:    nil,
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	dto := usecase.DeleteAccountCodesInputDTO{
		Id: idInt,
	}

	deleteAccountUsecase := usecase.NewDeleteAccountCodesUseCase(c.AccountCodesRepository)
	err = deleteAccountUsecase.Execute(dto)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "Nenhuma conta encontrada com o ID fornecido" {
			status = http.StatusNotFound
		}

		response := APIResponse{
			Message: "Erro ao deletar conta: " + err.Error(),
			Data:    nil,
		}
		w.WriteHeader(status)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	response := APIResponse{
		Message: "Conta deletada com sucesso",
		Data:    nil,
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (c *AccountCodesController) Update(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	idInt, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response := APIResponse{
			Message: "Erro ao atualizar conta: " + err.Error(),
			Data:    nil,
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	var dto usecase.CreateAccountCodesInputDTO
	err = json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		response := APIResponse{
			Message: "Erro ao atualizar conta: " + err.Error(),
			Data:    nil,
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	updateAccountUseCase := usecase.NewUpdateAccountCodesUseCase(c.AccountCodesRepository)
	output, err := updateAccountUseCase.Execute(idInt, dto)
	if err != nil {
		response := APIResponse{
			Message: "Erro ao atualizar conta: " + err.Error(),
			Data:    nil,
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	response := APIResponse{
		Message: "Conta atualizada com sucesso",
		Data:    output,
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
