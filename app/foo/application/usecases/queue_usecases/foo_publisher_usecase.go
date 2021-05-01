package queue_usecases

import (
	"go-course/demo/app/shared/domain/events"
	"go-course/demo/app/shared/domain/repositories"
)

type fooPublisherUseCase struct {
	queue repositories.PublisherQueue
}

func NewFooQueueUseCase(queue repositories.PublisherQueue) *fooPublisherUseCase {
	return &fooPublisherUseCase{
		queue: queue,
	}
}
func (f *fooPublisherUseCase) Execute(message *events.Message) error {
	topic := "topic"
	err := f.queue.Publish(topic, message)
	if err != nil {
		return err
	}
	return nil
}
