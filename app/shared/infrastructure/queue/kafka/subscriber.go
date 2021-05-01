package kafka

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/segmentio/kafka-go"
	"go-course/demo/app/shared/domain/events"
	"go-course/demo/app/shared/domain/repositories"
	"go-course/demo/app/shared/log"
	"go-course/demo/app/shared/utils"
	"time"
)

type subscriber struct {
	groupIdKafka string
	brokers      []string
	dialer       *kafka.Dialer
	useCase      repositories.QueueUseCaseRepository
}

func NewKafkaSubscriber(
	useCase repositories.QueueUseCaseRepository,
	groupIdKafka string,
	dialer *kafka.Dialer,
	brokers ...string) *subscriber {
	return &subscriber{
		useCase:      useCase,
		groupIdKafka: groupIdKafka,
		brokers:      brokers,
		dialer:       dialer,
	}
}

func (s *subscriber) getKafkaReader(topic string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        s.brokers,
		GroupID:        s.groupIdKafka,
		Topic:          topic,
		MinBytes:       1e6,  // 1MB
		MaxBytes:       10e6, // 10MB
		CommitInterval: time.Second,
		Dialer:         s.dialer,
		// Logger:      kafka.LoggerFunc(log.Info),
		ErrorLogger: kafka.LoggerFunc(log.Error),
	})
}

func (s *subscriber) Subscribe(topic string) {
	if topic == "" {
		err := errors.New("topic empty")
		log.WithError(err).Info("topic cannot be an empty string")
		return
	}
	reader := s.getKafkaReader(topic)
	defer reader.Close()
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.WithError(err).Error("error reading kafka message")
			continue
		}
		var payload *events.Message
		err = json.Unmarshal(msg.Value, &payload)
		if err != nil {
			log.WithError(err).Error("error trying to unmarshall message")
			continue
		}
		err = s.useCase.Execute(payload)
		if err != nil {
			log.WithError(err).Error(
				"%s error executing sender use case, value is: [ %s ]",
				topic,
				utils.EntityToJson(payload))
			continue
		}
		err = reader.CommitMessages(context.Background(), msg)
		if err != nil {
			log.WithError(err).Fatal("error to commit kafka message")
		}
	}
}