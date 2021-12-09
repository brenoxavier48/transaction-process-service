package entity

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
	return nil
}
