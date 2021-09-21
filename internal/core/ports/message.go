package ports

type Queueing interface {
	Publish(topic string, data []byte) error
}