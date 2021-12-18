package process_transaction

import (
	"github.com/brenoxavier48/imersaofc-gateway/domain/entity"
	"github.com/brenoxavier48/imersaofc-gateway/domain/repository"
)

type ProcessTransaction struct {
	Repository repository.TransactionRepository
}

func NewProcessTransaction(repository repository.TransactionRepository) *ProcessTransaction {
	return &ProcessTransaction{
		Repository: repository,
	}
}

func (p *ProcessTransaction) getCreditCardFromInput(input TransactionInputDTO) (*entity.CreditCard, error) {
	creditCard, invalidCreditCard := entity.NewCreditCard(
		input.CreditCardNumber,
		input.CreditCardName,
		input.CreditCardExpirationMonth,
		input.CreditCardExpirationYear,
		input.CreditCardCvv,
	)
	return creditCard, invalidCreditCard
}

func (p *ProcessTransaction) rejectTransaction(transaction *entity.Transaction, errorType error) (TransactionOutputDTO, error) {
	err := p.Repository.Insert(
		transaction.ID,
		transaction.AccountID,
		transaction.Amount,
		entity.REJECTED,
		errorType.Error(),
	)
	if err != nil {
		return TransactionOutputDTO{}, err
	}
	outPut := TransactionOutputDTO{
		ID:           transaction.ID,
		Status:       entity.REJECTED,
		ErrorMessage: errorType.Error(),
	}
	return outPut, nil
}

func (p *ProcessTransaction) Execute(input TransactionInputDTO) (TransactionOutputDTO, error) {
	creditCard, invalidCreditCard := p.getCreditCardFromInput(input)

	transaction := entity.NewTransaction()
	transaction.ID = input.ID
	transaction.AccountID = input.AccountID
	transaction.Amount = input.Amount

	if invalidCreditCard != nil {
		return p.rejectTransaction(transaction, invalidCreditCard)
	}

	transaction.SetCreditCard(*creditCard)
	invalidTransaction := transaction.IsValid()

	if invalidTransaction != nil {
		return p.rejectTransaction(transaction, invalidTransaction)
	}

	return TransactionOutputDTO{}, nil
}
