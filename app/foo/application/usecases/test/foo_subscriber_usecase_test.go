package test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-course/demo/app/foo/application/usecases/queue_usecases"
	"go-course/demo/app/foo/domain/entities"
	"testing"
)
func (r *repository) Create(foo *entities.Foo) error {
	mocked := r.Called(foo)
	if mocked.Error(0) != nil {
		return mocked.Error(0)
	}
	return nil
}

func TestFooSubscriberUseCase_Execute(t *testing.T) {
	t.Parallel()
	t.Run("When it tries to consume message, it executes use case successfully", func(t *testing.T) {
		repository := new(repository)
		repository.
			On("Create", mock.Anything).
			Return(nil)
		useCase := queue_usecases.NewFooSubscriberUseCase(repository)
		err := useCase.Execute(MessageIncoming)
		assert.NoError(t, err)
		repository.AssertExpectations(t)
	})

	t.Run("When it tries to consume message, it fails", func(t *testing.T) {
		repository := new(repository)
		repository.
			On("Create", mock.Anything).
			Return(errors.New("fails"))
		useCase := queue_usecases.NewFooSubscriberUseCase(repository)
		err := useCase.Execute(MessageIncoming)
		assert.Error(t, err)
		repository.AssertExpectations(t)
	})
}
