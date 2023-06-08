package queues

import (
	"context"
	"crypto/tls"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
)

func Produce(k_userid, k_pwd string) {

	log.Printf("Producint to topic : ")

	mechanism, err := scram.Mechanism(scram.SHA256, k_userid, k_pwd)

	if err != nil {

		log.Fatalln(err)

	}

	dialer := &kafka.Dialer{

		SASLMechanism: mechanism,

		TLS: &tls.Config{},
	}

	w := kafka.NewWriter(kafka.WriterConfig{

		Brokers: []string{"cool-werewolf-13410-eu1-kafka.upstash.io:9092"},

		//Topic: "$TOPIC",
		Topic: "inquiry",

		Dialer: dialer,
	})

	err = w.WriteMessages(context.Background(),
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}
}
