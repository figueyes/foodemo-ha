package config

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go-course/demo/app/shared/domain/constants"
	"go-course/demo/app/shared/log"
)

var connection *gorm.DB

type Connection interface {
	GetConnection() (*gorm.DB, error)
	CloseConnection()
}

type DbConnection struct {
	opts *Options
	url  string
}

func NewPostgresqlConnection(options ...*Options) *DbConnection {
	databaseOptions := MergeOptions(options...)
	url := databaseOptions.GetUrlConnection()
	if url == "" {
		log.Fatal(errors.New("error creating connection, url cannot be empty").Error())
	}
	return &DbConnection{
		opts: databaseOptions,
		url:  url,
	}
}

func (r *DbConnection) GetConnection() (*gorm.DB, error) {
	var err error
	if connection == nil || !isAlive() {
		log.Info("Trying connection to database")
		connection, err = gorm.Open(constants.POSTGRESQL, r.url)
		if err != nil {
			log.WithError(err).Error("Error trying connection to database")
			return nil, err
		} else {
			log.Info("Connection successfully")
		}
	}
	connection.LogMode(constants.LOGMODE)
	return connection, nil
}

func (r *DbConnection) CloseConnection() {
	if err := connection.Close(); err != nil {
		log.WithError(err).Error("Error trying to close connection")
	} else {
		log.Info("Connection was closed successfully")
	}
}

func isAlive() bool {
	if err := connection.DB().Ping(); err != nil {
		log.WithError(err).Error("Error trying to ping database")
		return false
	}
	return true
}
