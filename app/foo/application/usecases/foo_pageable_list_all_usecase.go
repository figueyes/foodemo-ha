package usecases

import (
	"go-course/demo/app/shared/domain/shared"
	"go-course/demo/app/shared/domain/repositories"
	"go-course/demo/app/shared/log"
)

type FooPageableListAllUseCase interface {
	ListPageable(limit, skip int64, query interface{}) (*shared.Pageable, error)
}

type fooPageableListAllUseCase struct {
	repository repositories.PageableRepository
}

func NewFooPageableListAllUseCase(repository repositories.PageableRepository) *fooPageableListAllUseCase {
	return &fooPageableListAllUseCase{
		repository: repository,
	}
}

func (f *fooPageableListAllUseCase) ListPageable(limit, skip int64, query interface{}) (*shared.Pageable, error) {
	list, err := f.repository.FindPageable(limit, skip, query)
	if err != nil {
		log.WithError(err).Info("error trying to find in database")
		return nil, err
	}
	return list, nil
}
