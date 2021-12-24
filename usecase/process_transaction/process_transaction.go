package process_transaction

import (
	"github.com/brenoxavier48/imersaofc-gateway/domain/entity"
	"github.com/brenoxavier48/imersaofc-gateway/domain/repository"
	"github.com/brenoxavier48/imersaofc-gateway/infra/broker"
)

type ProcessTransaction struct {
	Repository repository.TransactionRepository
	Producer   broker.Producer
	Topic      string
}

func NewProcessTransaction(repository repository.TransactionRepository, producer broker.Producer, topic string) *ProcessTransaction {
	return &ProcessTransaction{
		Repository: repository,
		Producer:   producer,
		Topic:      topic,
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
	p.producerPublish(outPut, []byte(transaction.ID))

	return outPut, nil
}

func (p *ProcessTransaction) acceptTransaction(transaction *entity.Transaction) (TransactionOutputDTO, error) {
	err := p.Repository.Insert(
		transaction.ID,
		transaction.AccountID,
		transaction.Amount,
		entity.ACCEPTED,
		"",
	)
	if err != nil {
		return TransactionOutputDTO{}, err
	}
	outPut := TransactionOutputDTO{
		ID:           transaction.ID,
		Status:       entity.ACCEPTED,
		ErrorMessage: "",
	}
	p.producerPublish(outPut, []byte(transaction.ID))

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

	return p.acceptTransaction(transaction)
}

func (p *ProcessTransaction) producerPublish(output TransactionOutputDTO, key []byte) error {
	err := p.Producer.Publish(output, key, p.Topic)
	if err != nil {
		return err
	}
	return nil
}
