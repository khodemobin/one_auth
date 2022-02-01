package messager

type Messager interface {
	Write(message, topic string) error
	Consumer(topic string, callback func(msg interface{})) error
}
