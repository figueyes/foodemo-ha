package config

import (
	"fmt"
	"go-course/demo/app/shared/domain/constants"
	"go-course/demo/app/shared/log"
	"os"
)

type Options struct {
	databaseName *string
	host         *string
	port         *int
	user         *string
	password     *string
}

func Config() *Options {
	return new(Options)
}

func (o *Options) DatabaseName(databaseName string) *Options {
	o.databaseName = &databaseName
	return o
}

func (o *Options) Host(host string) *Options {
	o.host = &host
	return o
}

func (o *Options) Port(port int) *Options {
	o.port = &port
	return o
}

func (o *Options) User(user string) *Options {
	o.user = &user
	return o
}

func (o *Options) Password(password string) *Options {
	o.password = &password
	return o
}

func MergeOptions(options ...*Options) *Options {
	option := new(Options)

	for _, v := range options {
		if v.databaseName != nil {
			option.databaseName = v.databaseName
		}
		if v.host != nil {
			option.host = v.host
		}
		if v.port != nil {
			option.port = v.port
		}
		if v.user != nil {
			option.user = v.user
		}
		if v.password != nil {
			option.password = v.password
		}
	}
	return option
}

var (
	defaultPort = 5432
)

func (o *Options) GetUrlConnection() string {
	UrlCockroachFormat := "postgresql://%v:%v@%v:%v/%v"

	if o.port == nil {
		o.port = &defaultPort
	}
	environment := os.Getenv(constants.ENVIRONMENT_TYPE)
	if environment == "local" || environment == "" {
		UrlCockroachFormat = "postgresql://%v:%v@%v:%v/%v?sslmode=disable"
	}
	log.Info("Connection: %s", fmt.Sprintf(UrlCockroachFormat, *o.user, "************", *o.host, *o.port, *o.databaseName))
	return fmt.Sprintf(UrlCockroachFormat, *o.user, *o.password, *o.host, *o.port, *o.databaseName)
}
