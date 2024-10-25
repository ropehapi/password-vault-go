package database

import (
	"database/sql"
	"fmt"
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

func (r *AccountRepository) GetByName(name string) ([]*entity.Account, error) {
	var accounts []*entity.Account

	rows, err := r.DB.Query("SELECT id, name, login, password, created_at, updated_at FROM accounts WHERE name LIKE ?", "%"+name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var account entity.Account
		err := rows.Scan(&account.ID, &account.Name, &account.Login, &account.Password, &account.CreatedAt, &account.UpdatedAt)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, &account)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return accounts, nil
}

func (r *AccountRepository) GetAll() ([]*entity.Account, error) {
	var accounts []*entity.Account

	rows, err := r.DB.Query("SELECT id, name, login, password, created_at, updated_at FROM accounts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var account entity.Account
		err := rows.Scan(&account.ID, &account.Name, &account.Login, &account.Password, &account.CreatedAt, &account.UpdatedAt)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, &account)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return accounts, nil
}

func (r *AccountRepository) Delete(id int64) error {
	stmt, err := r.DB.Prepare("DELETE FROM accounts where id=?")
	if err != nil {
		return err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("Nenhuma conta encontrada com o ID fornecido")
	}

	return nil
}
