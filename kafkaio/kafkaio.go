package kafkaio

import (
	"context"
	jsoniter "github.com/json-iterator/go"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
	"sync"
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

func Close(r *kafka.Reader, w *kafka.Writer) {
	log.Println("Closing kafka Reader...")
	err := r.Close()
	if err != nil {
		log.Println("Failed closing kafka Reader: ", err.Error())
	} else {
		log.Println("Closed kafka Reader...")
	}

	log.Println("Closing kafka Write...")
	err = w.Close()
	if err != nil {
		log.Println("Failed closing kafka Write: ", err.Error())
	} else {
		log.Println("Closed kafka Writer...")
	}
}

func ReceiveMessages[T any](ctx context.Context, wg *sync.WaitGroup, r *kafka.Reader, outputChannel chan<- T) {
	defer wg.Done()

	for processing := true; processing; {
		select {
		case <-ctx.Done():
			log.Println("Done with receiving messages...")
			close(outputChannel)
			processing = false
		default:
			payload, err := r.ReadMessage(ctx)
			if err != nil {
				log.Println("Unable to read message: ", err.Error())
				continue
			}

			var msg T
			err = jsoniter.Unmarshal(payload.Value, &msg)
			if err != nil {
				log.Println("Unable to parse message: ", err.Error())
				continue
			}

			outputChannel <- msg
		}
	}
}

func SendMessages[T any](ctx context.Context, wg *sync.WaitGroup, w *kafka.Writer, inputChannel <-chan T) {
	defer wg.Done()

	for msg := range inputChannel {
		json, err := jsoniter.MarshalIndent(msg, "", "  ")
		if err != nil {
			log.Println("Unable to convert to JSON: ", err.Error())
			continue
		}

		err = w.WriteMessages(ctx, kafka.Message{
			Value: json,
		})

		if err != nil {
			log.Println("Unable to send a message: ", err.Error())
		}
	}
}
