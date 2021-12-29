package main

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/brenoxavier48/imersaofc-gateway/infra/broker/kafka"
	"github.com/brenoxavier48/imersaofc-gateway/infra/factory"
	"github.com/brenoxavier48/imersaofc-gateway/infra/presenter/transaction"
	"github.com/brenoxavier48/imersaofc-gateway/usecase/process_transaction"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}

	repositoryFactory := factory.NewRepositoryDatabaseFactory(db)
	repository := repositoryFactory.CreateTransactionRepository()
	configMapProducer := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
	}
	kafkaPresenter := transaction.NewTransactionKafkaPresenter()
	kafkaProducer := kafka.NewKafkaProducer(configMapProducer, kafkaPresenter)

	var msgChan = make(chan *ckafka.Message)
	configMapConsumer := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"client.id":         "goapp",
		"group.id":          "goapp",
	}
	topics := []string{"transactions"}
	kafkaConsumer := kafka.NewConsumer(configMapConsumer, topics)

	go kafkaConsumer.Consume(msgChan)

	usecase := process_transaction.NewProcessTransaction(repository, kafkaProducer, "transactions_result")

	for msg := range msgChan {
		var input process_transaction.TransactionInputDTO
		json.Unmarshal(msg.Value, &input)
		usecase.Execute(input)
	}

}
