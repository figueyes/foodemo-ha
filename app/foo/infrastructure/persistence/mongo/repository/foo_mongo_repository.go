package repository

import (
	"go-course/demo/app/foo/domain/entities"
	"go-course/demo/app/foo/infrastructure/persistence/mongo/model"
	repository2 "go-course/demo/app/shared/infrastructure/persistence/mongo/repository"
	"go-course/demo/app/shared/log"
	"go-course/demo/app/shared/utils"
)

type FooMongoRepository struct {
	repository *repository2.MongoRepository
}

func NewFooMongoRepository(repository *repository2.MongoRepository) *FooMongoRepository {
	return &FooMongoRepository{
		repository: repository,
	}
}

func (f *FooMongoRepository) Create(foo *entities.Foo) error {
	log.Info("saving foo in database")
	utils.EntityToJson(foo)
	fooModel := new(*model.FooMongoModel)
	utils.ConvertEntity(foo, &fooModel)
	id, err := f.repository.Save(fooModel)
	if err != nil {
		log.WithError(err).Info("error trying to save in database")
		return err
	}
	log.Info("saved successfully with id %s", id)
	return nil
}

func (f *FooMongoRepository) Find() ([]*entities.Foo, error) {
	response, err := f.repository.Find(map[string]interface{}{})
	if err != nil {
		log.WithError(err).Info("error trying to find a element in database")
		return nil, err
	}
	foo := make([]*entities.Foo,0)
	utils.ConvertEntity(response,&foo)
	log.Info("%d foo was found: ", len(foo))
	utils.EntityToJson(foo)
	return foo, nil
}