package repository

import (
	"go-course/demo/app/foo/domain/entities"
	"go-course/demo/app/foo/infrastructure/persistence/mongo/model"
	shared "go-course/demo/app/shared/domain/shared"
	"go-course/demo/app/shared/infrastructure/persistence/mongo/repository"
	"go-course/demo/app/shared/log"
	"go-course/demo/app/shared/utils"
)

type FooMongoRepository struct {
	repository *repository.MongoRepository
}

func NewFooMongoRepository(repository *repository.MongoRepository) *FooMongoRepository {
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

func (f *FooMongoRepository) FindPageable(limit, page int64, query interface{}) (*shared.Pageable, error) {
	response, err := f.repository.FindPageable(limit, page, map[string]interface{}{})
	if err != nil {
		log.WithError(err).Info("error trying to find a element in database")
		return nil, err
	}
	foo := make([]*entities.Foo, 0)
	utils.ConvertEntity(response,&foo)
	log.Info("%d foo was found: ", len(foo))
	utils.EntityToJson(foo)
	pageable := &shared.Pageable{
		Total: len(foo),
		Page:  page,
		Limit: limit,
		Data:  foo,
	}
	return pageable, nil
}