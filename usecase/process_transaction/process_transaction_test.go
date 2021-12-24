package process_transaction

import (
	"testing"
	"time"

	"github.com/brenoxavier48/imersaofc-gateway/domain/entity"
	mock_repository "github.com/brenoxavier48/imersaofc-gateway/domain/repository/mock"
	mock_producer "github.com/brenoxavier48/imersaofc-gateway/infra/broker/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func makeSut(t *testing.T) (ProcessTransaction, *gomock.Controller, *mock_repository.MockTransactionRepository, *mock_producer.MockProducer) {
	ctrl := gomock.NewController(t)
	producerMock := mock_producer.NewMockProducer(ctrl)
	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	processTransaction := NewProcessTransaction(repositoryMock, producerMock, "transactions_result")
	return *processTransaction, ctrl, repositoryMock, producerMock
}

func assertHelper(t *testing.T, transactionInput TransactionInputDTO, transactionOutput TransactionOutputDTO) {
	processTransaction, ctrl, repositoryMock, producerMock := makeSut(t)
	defer ctrl.Finish()

	repositoryMock.EXPECT().
		Insert(transactionInput.ID, transactionInput.AccountID, transactionInput.Amount, transactionOutput.Status, transactionOutput.ErrorMessage).
		Return(nil)

	producerMock.EXPECT().
		Publish(transactionOutput, []byte(transactionOutput.ID), "transactions_result").
		Return(nil)

	output, err := processTransaction.Execute(transactionInput)

	assert.Nil(t, err)
	assert.Equal(t, transactionOutput, output)
}

func TestProcessTransaction_ExecuteInvalidCreditCard(t *testing.T) {
	invalidTransactionInput := TransactionInputDTO{
		ID:                        "1",
		AccountID:                 "1",
		CreditCardNumber:          "4000",
		CreditCardName:            "Any_Name",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  time.Now().Year(),
		CreditCardCvv:             123,
		Amount:                    100,
	}

	expectedTransactionOutPut := TransactionOutputDTO{
		ID:           "1",
		Status:       entity.REJECTED,
		ErrorMessage: "invalid credit card number",
	}
	assertHelper(t, invalidTransactionInput, expectedTransactionOutPut)
}

func TestProcessTransaction_ExecuteInvalidTransaction(t *testing.T) {
	invalidTransactionInput := TransactionInputDTO{
		ID:                        "1",
		AccountID:                 "1",
		CreditCardNumber:          "4193523830170205",
		CreditCardName:            "Any_Name",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  time.Now().Year(),
		CreditCardCvv:             123,
		Amount:                    10002,
	}

	expectedTransactionOutPut := TransactionOutputDTO{
		ID:           "1",
		Status:       entity.REJECTED,
		ErrorMessage: "you don't have enough limit for this transaction",
	}

	assertHelper(t, invalidTransactionInput, expectedTransactionOutPut)
}
func TestProcessTransaction_ExecuteValidTransaction(t *testing.T) {
	invalidTransactionInput := TransactionInputDTO{
		ID:                        "1",
		AccountID:                 "1",
		CreditCardNumber:          "4193523830170205",
		CreditCardName:            "Any_Name",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  time.Now().Year(),
		CreditCardCvv:             123,
		Amount:                    100,
	}

	expectedTransactionOutPut := TransactionOutputDTO{
		ID:           "1",
		Status:       entity.ACCEPTED,
		ErrorMessage: "",
	}

	assertHelper(t, invalidTransactionInput, expectedTransactionOutPut)
}
