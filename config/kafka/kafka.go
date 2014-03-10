package kafka

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/codegangsta/martini"
	"os"
	"path/filepath"
	"time"
)

type KafkaConfig struct {
	Hosts           []string "hosts"
	MetadataRetries int      "metadata_retries"
	WaitForElection int      "wait_for_election"
}

type Config map[string]KafkaConfig

func Kafka() martini.Handler {
	jsonFile, err := filepath.Abs("config/kafka/kafka.json")
	if err != nil {
		panic(err)
	}

	file, err := os.Open(jsonFile)
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(file)
	config := &Config{}

	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}

	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "development"
	}

	fmt.Printf("Using kafka config: %+v\n", (*config)[env])

	client, err := sarama.NewClient("swig", []string{"localhost:9092"}, &sarama.ClientConfig{MetadataRetries: 1, WaitForElection: 250 * time.Millisecond})

	if err != nil {
		panic(err)
	}

	defer client.Close()

	producer, err := sarama.NewProducer(client, &sarama.ProducerConfig{RequiredAcks: sarama.WaitForLocal, MaxBufferedBytes: 1, MaxBufferTime: 1})

	if err != nil {
		panic(err)
	}

	defer producer.Close()

	return func(context martini.Context) {
		context.Map(producer)
		context.Next()
	}
}
