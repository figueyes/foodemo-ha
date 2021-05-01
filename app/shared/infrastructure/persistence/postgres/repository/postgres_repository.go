package repository

import (
	"github.com/jinzhu/gorm"
	"go-course/demo/app/shared/domain/exceptions"
	"go-course/demo/app/shared/infrastructure/persistence/postgres/config"
	"go-course/demo/app/shared/log"
)

type PostgresRepository struct {
	connection config.Connection
}

func NewPostgresRepository(connection config.Connection) *PostgresRepository {
	return &PostgresRepository{
		connection: connection,
	}
}

func (p *PostgresRepository) Run() (*gorm.DB, error) {
	db, err := p.connection.GetConnection()
	if err != nil {
		log.Fatal("error connecting to postgres database")
		return nil, exceptions.ErrorConnectionDB
	}
	return db, nil
}
