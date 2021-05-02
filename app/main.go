package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go-course/demo/app/foo/application/usecases"
	"go-course/demo/app/foo/application/usecases/queue_usecases"
	"go-course/demo/app/foo/infrastructure/controllers"
	"go-course/demo/app/foo/infrastructure/persistence/mongo/repository"
	"go-course/demo/app/shared/domain/constants"
	"go-course/demo/app/shared/infrastructure/persistence/mongo"
	mongoRepository "go-course/demo/app/shared/infrastructure/persistence/mongo/repository"
	"go-course/demo/app/shared/infrastructure/queue/kafka"
	"go-course/demo/app/shared/infrastructure/queue/kafka/config"
	"go-course/demo/app/shared/log"
	"go-course/demo/app/version"
	"net/http"
	"os"
	"time"
)

var (
	_version      = "1.0.0"
	_groupId      = os.Getenv(constants.KAFKA_GROUP_ID)
	_kafkaBrokers = os.Getenv(constants.KAFKA_BROKERS)
)
// @title Foo API
// @version 1.0
// @description API for Golang Project Foo.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email aofiguer@uc.cl

// @license.name MIT
// @license.url http://google.com

// @BasePath /api/v1
func main() {
	if os.Getenv(constants.ENVIRONMENT_TYPE) == "local" {
		//utils.GetEnvironments()
	}
	log.Info("starting app %s", constants.APP)
	echoServer := echo.New()
	echoServer.Use(echoMiddleware.Recover())
	echoServer.Use(echoMiddleware.CORS())
	echoServer.HideBanner = true

	// mongo
	connection := mongo.CreateDbConnection()
	fooMongo := mongoRepository.NewMongoRepository("foo", connection)
	fooRepository := repository.NewFooMongoRepository(fooMongo)

	// postgres
	//connection := postgres.CreateDbConnection()
	//fooPostgres := postgresRepository.NewPostgresRepository(connection)
	//fooRepository := repository.NewFooPostgresRepository(fooPostgres)
	//postgres.AutomigratesEntities(connection)

	fooUseCase := queue_usecases.NewFooSubscriberUseCase(fooRepository)
	fooListAllUseCase := usecases.NewFooListAllUseCase(fooRepository)
	fooPageableListAllUseCase := usecases.NewFooPageableListAllUseCase(fooRepository)
	fooSubscriber := kafka.NewKafkaSubscriber(fooUseCase, _groupId, kafka_dialer.GetDialer(), _kafkaBrokers)
	log.Info("listening queues...")
	go fooSubscriber.Subscribe(os.Getenv(constants.TOPIC_DEMO))

	controllers.NewFooHandler(echoServer, fooListAllUseCase, fooPageableListAllUseCase)
	version.NewHealthHandler(echoServer, _version)

	echoServer.GET("/swagger/*", echoSwagger.WrapHandler)

	log.Info("Starting echoServer...")
	portServer := os.Getenv(constants.PORT_SERVER)
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", portServer),
		ReadTimeout:  3 * time.Minute,
		WriteTimeout: 3 * time.Minute,
	}

	echoServer.Logger.Fatal(echoServer.StartServer(server))
}
