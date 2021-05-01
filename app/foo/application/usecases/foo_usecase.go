package usecases

import (
	"go-course/demo/app/foo/domain/entities"
	"go-course/demo/app/foo/domain/repositories"
)

type FooListAllUseCase interface {
	List() ([]*entities.Foo, error)
}

type fooListAllUseCase struct {
	repository repositories.FooFindRepository
}

func NewFooListAllUseCase(repository repositories.FooFindRepository) *fooListAllUseCase {
	return &fooListAllUseCase{
		repository: repository,
	}
}

func (f *fooListAllUseCase) List() ([]*entities.Foo, error) {
	result, err := f.repository.Find()
	if err != nil {
		return nil, err
	}
	return result, nil
}
