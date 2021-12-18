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

func (p *ProcessTransaction) Execute(input TransactionInputDTO) (TransactionOutputDTO, error) {
	_, invalidCreditCard := p.getCreditCardFromInput(input)

	transaction := entity.NewTransaction()
	transaction.ID = input.ID
	transaction.AccountID = input.AccountID
	transaction.Amount = input.Amount

	if invalidCreditCard != nil {
		p.Repository.Insert(
			transaction.ID,
			transaction.AccountID,
			transaction.Amount,
			entity.REJECTED,
			invalidCreditCard.Error(),
		)
		outPut := TransactionOutputDTO{
			ID:           transaction.ID,
			Status:       entity.REJECTED,
			ErrorMessage: invalidCreditCard.Error(),
		}
		return outPut, nil
	}

	return TransactionOutputDTO{}, nil
}
