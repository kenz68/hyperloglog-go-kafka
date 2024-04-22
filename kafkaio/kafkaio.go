package kafkaio

import (
	"github.com/segmentio/kafka-go"
	"os"
)

func Init() (*kafka.Reader, *kafka.Writer) {
	kBroker, set := os.LookupEnv("KAFKA_BROKER")
	if !set {
		kBroker = "localhost:9092"
	}

	inTopic, set := os.LookupEnv("USERS_TOPIC")
	if !set {
		inTopic = "users"
	}

	outTopic, set := os.LookupEnv("STATS_TOPIC")
	if !set {
		outTopic = "stats"
	}

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kBroker},
		Topic:   inTopic,
	})

	//https://pkg.go.dev/github.com/segmentio/kafka-go#Writer
	w := &kafka.Writer{
		Addr:  kafka.TCP(kBroker),
		Topic: outTopic,
	}
	return r, w
}
