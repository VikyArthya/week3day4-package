package finance

import "errors"

type Account struct {
	Name  string
	Email string
	Saldo Saldo
}

func NewAccount(name string, email string) (*Account, error) {
	if name == "" || email == "" {
		return nil, errors.New("nama dan email tidak boleh kosong")
	}

	return &Account{
		Name:  name,
		Email: email,
		Saldo: Saldo{Amount: 0}, // saldo awal adalah 0
	}, nil
}
