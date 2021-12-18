package process_transaction

import (
	"testing"
	"time"

	"github.com/brenoxavier48/imersaofc-gateway/domain/entity"
	mock_repository "github.com/brenoxavier48/imersaofc-gateway/domain/repository/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

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

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().
		Insert(invalidTransactionInput.ID, invalidTransactionInput.AccountID, invalidTransactionInput.Amount, expectedTransactionOutPut.Status, expectedTransactionOutPut.ErrorMessage).
		Return(nil)

	processTransaction := NewProcessTransaction(repositoryMock)

	output, err := processTransaction.Execute(invalidTransactionInput)

	assert.Nil(t, err)
	assert.Equal(t, expectedTransactionOutPut, output)
}
