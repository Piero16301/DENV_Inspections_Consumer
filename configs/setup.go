package configs

import (
	"fmt"
	"github.com/IBM/sarama"
	"log"
)

func ConnectConsumer() sarama.Consumer {
	fmt.Println("Conectando a Kafka...")

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	kfProperties, err := GetKafkaProperties()
	if err != nil {
		log.Fatal(err)
	}

	consumer, err := sarama.NewConsumer([]string{kfProperties.GetDSN()}, config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conectado a Kafka en " + kfProperties.GetDSN())

	return consumer
}

// Consumer Instancia de Cliente
var Consumer = ConnectConsumer()

// BaseUrl URL Base del Backend
var BaseUrl = "http://127.0.0.1:8080"
