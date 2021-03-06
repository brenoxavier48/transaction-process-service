package entity

import "errors"

const (
	REJECTED = "rejected"
	ACCEPTED = "accepted"
)

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
	} else if t.Amount < 1 {
		return errors.New("the amount must be greater than 1")
	}
	return nil
}

func (t *Transaction) SetCreditCard(creditCard CreditCard) {
	t.CreditCard = creditCard
}
