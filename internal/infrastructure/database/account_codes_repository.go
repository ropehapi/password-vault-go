package database

import (
	"database/sql"
	"fmt"
	"github.com/ropehapi/password-vault-go/internal/domain/entity"
	"strconv"
)

type AccountCodesRepository struct {
	DB *sql.DB
}

func NewAccountCodesRepository(db *sql.DB) *AccountCodesRepository {
	return &AccountCodesRepository{
		DB: db,
	}
}

func (r *AccountCodesRepository) Save(accountCode *entity.AccountCodes) error {
	stmt, err := r.DB.Prepare("INSERT INTO account_codes (name, codes) values(?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(accountCode.Name, accountCode.Codes)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	accountCode.ID = strconv.FormatInt(id, 10)

	return nil
}

func (r *AccountCodesRepository) GetByName(name string) ([]*entity.AccountCodes, error) {
	var accountsCodes []*entity.AccountCodes

	rows, err := r.DB.Query("SELECT id, name, codes, created_at, updated_at FROM account_codes WHERE name LIKE ?", "%"+name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var accountCodes entity.AccountCodes
		err := rows.Scan(&accountCodes.ID, &accountCodes.Name, &accountCodes.Codes, &accountCodes.CreatedAt, &accountCodes.UpdatedAt)
		if err != nil {
			return nil, err
		}

		accountsCodes = append(accountsCodes, &accountCodes)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return accountsCodes, nil
}

func (r *AccountCodesRepository) GetAll() ([]*entity.AccountCodes, error) {
	var accountsCodes []*entity.AccountCodes

	rows, err := r.DB.Query("SELECT id, name, codes, created_at, updated_at FROM account_codes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var accountCodes entity.AccountCodes
		err := rows.Scan(&accountCodes.ID, &accountCodes.Name, &accountCodes.Codes, &accountCodes.CreatedAt, &accountCodes.UpdatedAt)
		if err != nil {
			return nil, err
		}

		accountsCodes = append(accountsCodes, &accountCodes)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return accountsCodes, nil
}

func (r *AccountCodesRepository) Delete(id int64) error {
	stmt, err := r.DB.Prepare("DELETE FROM account_codes where id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

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

func (r *AccountCodesRepository) Update(id int64, accountCodes *entity.AccountCodes) error {
	stmt, err := r.DB.Prepare("UPDATE account_codes SET name = ?, codes = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(accountCodes.Name, accountCodes.Codes, id)
	if err != nil {
		return err
	}

	accountCodes.ID = strconv.FormatInt(id, 10)

	return nil
}
