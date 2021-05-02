package test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"go-course/demo/app/foo/application/usecases"
	"go-course/demo/app/foo/domain/entities"
	"go-course/demo/app/shared/utils"
	"testing"
)

func (r *repository) Find() ([]*entities.Foo, error) {
	mocked := r.Called()
	if mocked.Get(0) == nil {
		return nil, mocked.Error(1)
	}
	response := make([]*entities.Foo, 0)
	utils.ConvertEntity(mocked.Get(0), &response)
	return response, nil
}

func TestFooListAllUseCase_List(t *testing.T) {
	t.Parallel()
	t.Run("When it tries to list foo array, it shows foo array", func(t *testing.T) {
		repository := new(repository)
		repository.
			On("Find").
			Return(FooArray, nil)
		useCase := usecases.NewFooListAllUseCase(repository)
		_, err := useCase.List()
		assert.NoError(t, err)
		repository.AssertExpectations(t)
	})
	t.Run("When it tries to list foo array, it fails", func(t *testing.T) {
		repository := new(repository)
		repository.
			On("Find").
			Return(nil, errors.New("test error"))
		useCase := usecases.NewFooListAllUseCase(repository)
		_, err := useCase.List()
		assert.Error(t, err)
		repository.AssertExpectations(t)
	})
}
