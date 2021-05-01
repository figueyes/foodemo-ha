package queue_usecases

import (
	"github.com/google/uuid"
	"go-course/demo/app/foo/domain/entities"
	"go-course/demo/app/foo/domain/repositories"
	"go-course/demo/app/shared/domain/events"
	"go-course/demo/app/shared/utils"
)

type fooSubscriberUseCase struct {
	repository repositories.FooCreateRepository
}

func NewFooSubscriberUseCase(repository repositories.FooCreateRepository) *fooSubscriberUseCase {
	return &fooSubscriberUseCase{
		repository: repository,
	}
}

func (f *fooSubscriberUseCase) Execute(message *events.Message) error {
	payload := events.PayloadFoo{}
	utils.ConvertEntity(message.Payload, &payload)
	foo := &entities.Foo{
		Id:      uuid.New().String(),
		Message: payload.Message,
	}
	err := f.repository.Create(foo)
	if err != nil {
		return err
	}
	return nil
}
