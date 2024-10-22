package entity

type AccountRepositoryInterface interface {
	Save(account *Account) error
	FindByName(id string) (*Account, error)
}
