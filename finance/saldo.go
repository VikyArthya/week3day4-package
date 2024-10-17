package finance

import "errors"

type Saldo struct {
	Amount int
}

func (s *Saldo) AddSaldo(amount int) error {
	if amount <= 0 {
		return errors.New("jumlah saldo yang ditambahkan harus lebih dari 0")
	}
	s.Amount += amount
	return nil
}

func (s *Saldo) DeductSaldo(amount int) error {
	if amount <= 0 {
		return errors.New("jumlah saldo yang dikurangi harus lebih dari 0")
	}
	if s.Amount < amount {
		return errors.New("saldo tidak mencukupi untuk pengurangan")
	}
	s.Amount -= amount
	return nil
}
