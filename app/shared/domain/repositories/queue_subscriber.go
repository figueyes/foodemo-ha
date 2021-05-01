package repositories

type SubscriberQueue interface {
	Subscribe(topic string)
}
