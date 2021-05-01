package mongo

import (
	"go-course/demo/app/shared/domain/constants"
	"go-course/demo/app/shared/infrastructure/persistence/mongo/config"
	"go-course/demo/app/shared/log"
	"os"
	"strconv"
)

func CreateDbConnection() *config.DbConnection {
	var port int
	var err error

	dbHost := os.Getenv(constants.MONGODB_HOST)
	dbPort := os.Getenv(constants.MONGODB_PORT)
	dbDatabase := os.Getenv(constants.MONGODB_DATABASE)
	dbUsername := os.Getenv(constants.MONGODB_USERNAME)
	dbPassword := os.Getenv(constants.MONGODB_PASSWORD)

	if len(dbHost) == 0 || len(dbPort) == 0 || len(dbDatabase) == 0 ||
		len(dbUsername) == 0 || len(dbPassword) == 0 {
		log.Fatal("data connection invalid")
	}
	if port, err = strconv.Atoi(dbPort); err != nil {
		log.Fatal("invalid port error")
	}
	connection := config.NewMongoConnection(config.Config().
		Host(dbHost).
		Port(port).
		DatabaseName(dbDatabase).
		User(dbUsername).
		Password(dbPassword),
	)
	log.Info("%s connection successfully", constants.MONGO)
	return connection
}