package test

import (
	"github.com/stretchr/testify/mock"
	"go-course/demo/app/foo/domain/entities"
	shared "go-course/demo/app/shared/domain/entities"
	"go-course/demo/app/shared/domain/events"
	"time"
)

type repository struct {
	mock.Mock
}

var (
	MessageIncoming = &events.Message{
		EventId:         "test id",
		EventName:       "test message",
		EventDataFormat: "json",
		Type:            "test",
		Timestamp:       time.Now().String(),
		Version:         "1.0.0",
		Country:         "cl",
		Origin:          "demo",
		Payload: map[string]interface{}{
			"message": "hello test",
		},
	}
	FooArray = []*entities.Foo{
		{
			Id:      "test",
			Message: "that's a testing message",
		},
	}
	FooPaged = &shared.Pageable{
		Total: 1,
		Page:  1,
		Limit: 10,
		Data:  FooArray,
	}
)
