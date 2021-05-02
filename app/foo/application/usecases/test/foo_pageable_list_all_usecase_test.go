package test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-course/demo/app/foo/application/usecases"
	"go-course/demo/app/shared/domain/entities"
	"go-course/demo/app/shared/utils"
	"testing"
)

func (r *repository) FindPageable(limit, page int64, query interface{}) (*entities.Pageable, error) {
	mocked := r.Called(limit, page, query)
	if mocked.Get(0) == nil {
		return nil, mocked.Error(1)
	}
	response := new(entities.Pageable)
	utils.ConvertEntity(mocked.Get(0), &response)
	return response, nil
}

func TestFooPageableListAllUseCase_ListPageable(t *testing.T) {
	t.Parallel()
	t.Run("When it tries to list foo pageable array, it shows foos", func(t *testing.T) {
		repository := new(repository)
		repository.
			On("FindPageable", mock.Anything, mock.Anything, mock.Anything).
			Return(FooPaged, nil)
		useCase := usecases.NewFooPageableListAllUseCase(repository)
		_, err := useCase.ListPageable(10, 1, map[string]interface{}{})
		assert.NoError(t, err)
		repository.AssertExpectations(t)
	})
	t.Run("When it tries to list foo pageable array, it fails", func(t *testing.T) {
		repository := new(repository)
		repository.
			On("FindPageable", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, errors.New("test error"))
		useCase := usecases.NewFooPageableListAllUseCase(repository)
		_, err := useCase.ListPageable(10, 1, map[string]interface{}{})
		assert.Error(t, err)
		repository.AssertExpectations(t)
	})
}
