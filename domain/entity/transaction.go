package entity

import "errors"

type Transaction struct {
	ID           string
	AccountID    string
	Amount       float64
	Status       string
	CreditCard   CreditCard
	ErrorMessage string
}

func NewTransaction() *Transaction {
	return &Transaction{}
}

func (t *Transaction) IsValid() error {
	if t.Amount > 1000 {
		return errors.New("you don't have enough limit for this transaction")
	}
	return nil
}
