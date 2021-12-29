package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreditCardNumber(t *testing.T) {
	_, err := NewCreditCard("9939393999999", "Jose Augusto", 12, 2026, 126)
	assert.Equal(t, "invalid credit card number", err.Error())

	_, err = NewCreditCard("4193523830170205", "Jose Augusto", 12, 2026, 126)
	assert.Nil(t, err)
}

func TestCreditCardExpirationMonth(t *testing.T) {
	_, err := NewCreditCard("4193523830170205", "Jose Augusto", 13, 2026, 126)
	assert.Equal(t, "invalid expiration month, must be less than 13", err.Error())

	_, err = NewCreditCard("4193523830170205", "Jose Augusto", 0, 2026, 126)
	assert.Equal(t, "invalid expiration month, must be more than 0", err.Error())

	_, err = NewCreditCard("4193523830170205", "Jose Augusto", 6, 2026, 126)
	assert.Nil(t, err)
}

func TestCreditCardExpirationYear(t *testing.T) {
	lastYear := time.Now().AddDate(-1, 0, 0)

	_, err := NewCreditCard("4193523830170205", "Jose Augusto", 10, lastYear.Year(), 126)
	assert.Equal(t, "invalid expiration year", err.Error())
}
