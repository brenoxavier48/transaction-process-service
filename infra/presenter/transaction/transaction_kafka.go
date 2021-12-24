package transaction

import (
	"encoding/json"
	"log"

	"github.com/brenoxavier48/imersaofc-gateway/usecase/process_transaction"
)

type TransactionKafkaPresenter struct {
	ID           string `json:"id"`
	Status       string `json:"status"`
	ErrorMessage string `json:"error_message"`
}

func NewTransactionKafkaPresenter() *TransactionKafkaPresenter {
	return &TransactionKafkaPresenter{}
}

func (t *TransactionKafkaPresenter) Bind(input interface{}) error {
	t.ID = input.(process_transaction.TransactionOutputDTO).ID
	t.Status = input.(process_transaction.TransactionOutputDTO).Status
	t.ErrorMessage = input.(process_transaction.TransactionOutputDTO).ErrorMessage

	return nil
}

func (t *TransactionKafkaPresenter) Show() ([]byte, error) {
	jsonCreated, err := json.Marshal(t)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return jsonCreated, nil
}
