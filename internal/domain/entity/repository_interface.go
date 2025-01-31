package entity

type AccountRepositoryInterface interface {
	Save(account *Account) error
	GetByName(name string) ([]*Account, error)
	GetAll() ([]*Account, error)
	Delete(id int64) error
	Update(id int64, account *Account) error
}

type AccountCodesRepositoryInterface interface {
	Save(account *Account) error
	GetByName(name string) ([]*Account, error)
	GetAll() ([]*Account, error)
	Delete(id int64) error
	Update(id int64, account *Account) error
}
