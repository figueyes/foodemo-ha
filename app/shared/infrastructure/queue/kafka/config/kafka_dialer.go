package kafka_dialer

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
	"go-course/demo/app/shared/log"
	"os"
)

func GetDialer() *kafka.Dialer {
	if os.Getenv("ENVIRONMENT_TYPE") == "local" || os.Getenv("ENVIRONMENT_TYPE") == "" {
		return kafka.DefaultDialer
	}

	kafkaUsername := os.Getenv("KAFKA_USERNAME")
	kafkaPassword := os.Getenv("KAFKA_PASSWORD")

	validUserNameAndPassword(kafkaUsername, kafkaPassword)

	rootCAs, _ := x509.SystemCertPool()

	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}
	dialer := &kafka.Dialer{
		DualStack: true,
		SASLMechanism: plain.Mechanism{
			Username: kafkaUsername, // access key
			Password: kafkaPassword, // secret
		},
		TLS: &tls.Config{
			InsecureSkipVerify: true,
			RootCAs:            rootCAs,
		},

	}
	return dialer
}

func validUserNameAndPassword(kafkaUsername string, kafkaPassword string) {
	if len(kafkaUsername) == 0 || len(kafkaPassword) == 0 {
		log.Fatal("username and password are required to connect to kafka")
	}
}
