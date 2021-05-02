package test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-course/demo/app/foo/application/usecases/queue_usecases"
	"testing"
)

func (r *repository) Publish(topic string, data interface{}) error {
	mocked := r.Called(topic, data)
	if mocked.Error(0) != nil {
		return mocked.Error(0)
	}
	return nil
}

func TestFooPublisherUseCase_Execute(t *testing.T) {
	t.Parallel()
	t.Run("When it tries to publish message, it sends to queue", func(t *testing.T) {
		repository := new(repository)
		repository.
			On("Publish", mock.Anything, mock.Anything).
			Return(nil)
		useCase := queue_usecases.NewFooQueueUseCase(repository)
		err := useCase.Execute(MessageIncoming)
		assert.NoError(t, err)
		repository.AssertExpectations(t)
	})

	t.Run("When it tries to publish message, it fails", func(t *testing.T) {
		repository := new(repository)
		repository.
			On("Publish", mock.Anything, mock.Anything).
			Return(errors.New("fails"))
		useCase := queue_usecases.NewFooQueueUseCase(repository)
		err := useCase.Execute(MessageIncoming)
		assert.Error(t, err)
		repository.AssertExpectations(t)
	})
}
