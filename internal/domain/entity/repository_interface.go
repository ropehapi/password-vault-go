package entity

type AccountRepositoryInterface interface {
	Save(account *Account) error
	FindByName(name string) (*Account, error)
}
