package repository

import (
	"github.com/jinzhu/gorm"
	"go-course/demo/app/foo/domain/entities"
	"go-course/demo/app/foo/infrastructure/persistence/postgres/model"
	"go-course/demo/app/shared/infrastructure/persistence/postgres/repository"
	"go-course/demo/app/shared/log"
	"go-course/demo/app/shared/utils"
)

type FooPostgresRepository struct {
	repository *repository.PostgresRepository
	db         *gorm.DB
}

func NewFooPostgresRepository(repository *repository.PostgresRepository) *FooPostgresRepository {
	db, err := repository.Run()
	if err != nil {
		panic(`cannot initialize db`)
	}
	fooPostgresRepository := &FooPostgresRepository{
		repository: repository,
		db: db,
	}
	return fooPostgresRepository
}

func (f *FooPostgresRepository) Create(foo *entities.Foo) error {
	log.Info("saving foo in database")
	utils.EntityToJson(foo)
	saver := new(model.FooPostgresModel)
	utils.ConvertEntity(foo, &saver)
	defer f.db.Close()
	err := f.db.Create(saver).Error
	if err != nil {
		return err
	}
	log.Info("saved successfully")
	return nil
}

func (f *FooPostgresRepository) Find() ([]*entities.Foo, error) {
	fooPostgresModel := make([]*model.FooPostgresModel, 0)
	defer f.db.Close()
	err := f.db.
		Model(&fooPostgresModel).
		Scan(&fooPostgresModel).Error
	if err != nil {
		return nil, err
	}
	foo := make([]*entities.Foo, 0)
	for _, value := range fooPostgresModel {
		f := &entities.Foo{
			Id:      value.Id,
			Message: value.Message,
		}
		foo = append(foo, f)
	}
	log.Info("%d foo was found: ", len(foo))
	utils.EntityToJson(foo)
	return foo, nil

}
