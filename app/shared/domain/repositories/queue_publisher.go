package repositories

type PublisherQueue interface {
	Publish(topic string, data interface{}) error
}
