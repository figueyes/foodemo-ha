package repositories

import "go-course/demo/app/shared/domain/events"

type QueueUseCaseRepository interface {
	Execute(message *events.Message) error
}
