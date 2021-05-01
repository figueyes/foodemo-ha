package kafka

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"
	"go-course/demo/app/shared/log"
	"go-course/demo/app/shared/utils"
	"reflect"
	"time"
)

type publisher struct {
	brokers      []string
	dialer       *kafka.Dialer
	kafkaWriters map[string]*kafka.Writer
}

func NewKafkaPublisher(kafkaDialer *kafka.Dialer, brokers ...string) *publisher {
	return &publisher{
		brokers:      brokers,
		kafkaWriters: make(map[string]*kafka.Writer, 0),
		dialer:       kafkaDialer,
	}
}

func (p *publisher) getKafkaWriter(topic string) *kafka.Writer {
	if p.kafkaWriters[topic] == nil {
		p.kafkaWriters[topic] = kafka.NewWriter(kafka.WriterConfig{
			Brokers:          p.brokers,
			Topic:            topic,
			Balancer:         &kafka.LeastBytes{},
			CompressionCodec: snappy.NewCompressionCodec(),
			BatchSize:        1,
			BatchTimeout:     10 * time.Millisecond,
			Dialer:           p.dialer,
		})
	}
	return p.kafkaWriters[topic]
}

func (p *publisher) Publish(topic string, data interface{}) error {
	kafkaMessages, err := createKafkaMessages(data)
	if err != nil {
		return errors.New("error on publishing kafka message")
	}
	kafkaWriter := p.getKafkaWriter(topic)
	err = kafkaWriter.WriteMessages(context.Background(), kafkaMessages...)
	if err != nil {
		log.WithError(err).Error("error writing kafka message")
		return err
	} else {
		log.WithField("topic", topic).Info("kafka message published successfully")
		log.Info("message bytes size: %d bytes", uint(p.kafkaWriters[topic].Stats().Bytes))
	}
	return nil
}

func createKafkaMessages(data interface{}) ([]kafka.Message, error) {
	var kafkaMessages []kafka.Message
	if utils.IsNilFixed(data) {
		return nil, errors.New("kafka error the data is empty")
	}
	switch reflect.TypeOf(data).Kind() {
	case reflect.Array, reflect.Slice:
		value := reflect.ValueOf(data)
		for index := 0; index < value.Len(); index++ {
			kafkaMessages = append(kafkaMessages, createMessageKafka(value.Index(index)))
		}
	default:
		kafkaMessages = append(kafkaMessages, createMessageKafka(data))
	}
	return kafkaMessages, nil
}

func createMessageKafka(data interface{}) kafka.Message {
	payload := utils.EntityToJson(data)
	key := uuid.New().String()
	kafkaMessage := kafka.Message{
		Key:   []byte(key),
		Value: []byte(payload),
	}
	return kafkaMessage
}

func (p *publisher) Close(topic string) error {
	if p.kafkaWriters[topic] == nil {
		return errors.New("error trying to close kafka connection, connection does not exist")
	}
	p.kafkaWriters[topic].Close()
	delete(p.kafkaWriters, topic)
	return nil
}
