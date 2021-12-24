package kafka

import (
	"testing"

	"github.com/brenoxavier48/imersaofc-gateway/domain/entity"
	"github.com/brenoxavier48/imersaofc-gateway/infra/presenter/transaction"
	"github.com/brenoxavier48/imersaofc-gateway/usecase/process_transaction"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/stretchr/testify/assert"
)

func TestProducerPublish(t *testing.T) {
	expectedTransactionOutPut := process_transaction.TransactionOutputDTO{
		ID:           "1",
		Status:       entity.REJECTED,
		ErrorMessage: "invalid credit card number",
	}
	configMap := ckafka.ConfigMap{
		"test.mock.num.brokers": 3,
	}
	producer := NewKafkaProducer(&configMap, transaction.NewTransactionKafkaPresenter())
	err := producer.Publish(expectedTransactionOutPut, []byte("1"), "test")
	assert.Nil(t, err)

}
