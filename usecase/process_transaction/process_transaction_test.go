package process_transaction

import (
	"testing"
	"time"

	"github.com/brenoxavier48/imersaofc-gateway/domain/entity"
	mock_repository "github.com/brenoxavier48/imersaofc-gateway/domain/repository/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func assertHelper(t *testing.T, transactionInput TransactionInputDTO, transactionOutput TransactionOutputDTO) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().
		Insert(transactionInput.ID, transactionInput.AccountID, transactionInput.Amount, transactionOutput.Status, transactionOutput.ErrorMessage).
		Return(nil)

	processTransaction := NewProcessTransaction(repositoryMock)

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
