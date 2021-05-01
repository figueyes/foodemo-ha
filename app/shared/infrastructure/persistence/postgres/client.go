package postgres

import (
	"go-course/demo/app/foo/infrastructure/persistence/postgres/model"
	"go-course/demo/app/shared/domain/constants"
	"go-course/demo/app/shared/infrastructure/persistence/postgres/config"
	"go-course/demo/app/shared/log"
	"os"
	"strconv"
)

func AutomigratesEntities(connection config.Connection) {
	log.Info("Migrating entities...")
	migrate := config.NewMigrate(connection)
	migrate.AutoMigrateAll(
		model.FooPostgresModel{},
	)
	log.Info("Migration finalised successfully")
}

func CreateDbConnection() *config.DbConnection {
	var port int
	var err error

	dbHost := os.Getenv(constants.POSTGRES_HOST)
	dbPort := os.Getenv(constants.POSTGRES_PORT)
	dbDatabase := os.Getenv(constants.POSTGRES_DATABASE)
	dbUsername := os.Getenv(constants.POSTGRES_USERNAME)
	dbPassword := os.Getenv(constants.POSTGRES_PASSWORD)

	if len(dbHost) == 0 || len(dbPort) == 0 || len(dbDatabase) == 0 ||
		len(dbUsername) == 0 || len(dbPassword) == 0 {
		log.Fatal("data connection invalid")
	}
	if port, err = strconv.Atoi(dbPort); err != nil {
		log.Fatal("invalid port error")
	}
	connection := config.NewPostgresqlConnection(config.Config().
		Host(dbHost).
		Port(port).
		DatabaseName(dbDatabase).
		User(dbUsername).
		Password(dbPassword),
	)
	return connection
}
