package entity

type AccountRepositoryInterface interface {
	Save(account *Account) error
	GetByName(name string) ([]*Account, error)
	GetAll() ([]*Account, error)
}
