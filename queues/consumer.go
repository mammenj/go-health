package queues

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
)

func Consume(k_userid, k_pwd string) {
	mechanism, err := scram.Mechanism(scram.SHA256, k_userid, k_pwd)

	if err != nil {

		log.Fatalln(err)

	}

	dialer := &kafka.Dialer{

		SASLMechanism: mechanism,

		TLS: &tls.Config{},
	}

	r := kafka.NewReader(kafka.ReaderConfig{

		Brokers: []string{"cool-werewolf-13410-eu1-kafka.upstash.io:9092"},

		//GroupID: "$GROUP_NAME",
		GroupID: "mygroup",

		//Topic: "$TOPIC",
		Topic: "inquiry",

		Dialer: dialer,
	})

	// ...
	message, err := r.ReadMessage(context.Background())
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Messages...")
	fmt.Println(string(message.Key), string(message.Value))

	r.Close()

}
