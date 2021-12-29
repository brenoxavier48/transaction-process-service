package kafka

import (
	"log"

	"github.com/brenoxavier48/imersaofc-gateway/infra/presenter"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaProducer struct {
	KafkaConfigMap *ckafka.ConfigMap
	presenter      presenter.Presenter
}

func NewKafkaProducer(KafkaConfigMap *ckafka.ConfigMap, presenter presenter.Presenter) *KafkaProducer {
	return &KafkaProducer{
		KafkaConfigMap: KafkaConfigMap,
		presenter:      presenter,
	}
}

func (p *KafkaProducer) Publish(msg interface{}, key []byte, topic string) error {
	producer, err := ckafka.NewProducer(p.KafkaConfigMap)
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = p.presenter.Bind(msg)
	if err != nil {
		log.Fatal(err)
		return err
	}

	presenterMsg, err := p.presenter.Show()
	if err != nil {
		log.Fatal(err)
		return err
	}

	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: int32(ckafka.PartitionAny)},
		Value:          presenterMsg,
		Key:            key,
	}
	err = producer.Produce(message, nil)
	if err != nil {
		panic(err)
	}
	return nil
}
