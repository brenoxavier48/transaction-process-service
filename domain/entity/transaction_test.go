package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransaction_IsValid(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 900
	assert.Nil(t, transaction.IsValid())
}

func TestTransaction_AmountLimit(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 1001
	err := transaction.IsValid()
	if assert.Error(t, err) {
		assert.Equal(t, "you don't have enough limit for this transaction", err.Error())
	}
}
