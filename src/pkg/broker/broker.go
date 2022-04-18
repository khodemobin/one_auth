package broker

type Broker interface {
	Write(message, topic string) error
	Consumer(topic string, callback func(msg interface{})) error
}
