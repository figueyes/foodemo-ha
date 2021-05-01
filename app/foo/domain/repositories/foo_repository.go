package repositories

import "go-course/demo/app/foo/domain/entities"

type FooCreateRepository interface {
	Create(foo *entities.Foo) error
}

type FooFindRepository interface {
	Find() ([]*entities.Foo, error)
}