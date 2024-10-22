package database

import (
	"database/sql"
	"github.com/ropehapi/password-vault-go/internal/domain/entity"
)

type AccountRepository struct {
	DB *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{
		DB: db,
	}
}

func (r *AccountRepository) Save(account *entity.Account) error {
	stmt, err := r.DB.Prepare("INSERT INTO accounts (name, login, password) values(?,?,?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(account.Name, account.Login, account.Password)
	if err != nil {
		return err
	}

	return nil
}

func (r *AccountRepository) FindByName(name string) (*entity.Account, error) {
	var account entity.Account

	err := r.DB.QueryRow("SELECT id, name, login, password, created_at, updated_at FROM account WHERE name = $1", name).Scan(&account.ID, &account.Name, &account.Login, &account.Password, &account.CreatedAt, &account.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &account, nil
}
