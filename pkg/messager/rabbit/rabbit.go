package rabbit

import (
	"errors"
	"fmt"

	"github.com/isayme/go-amqp-reconnect/rabbitmq"
	"github.com/khodemobin/pilo/auth/internal/config"
	"github.com/khodemobin/pilo/auth/pkg/logger"
	"github.com/khodemobin/pilo/auth/pkg/messager"
	"github.com/streadway/amqp"
)

const TopicProcess = "data_queue"

type rabbit struct {
	conn   *rabbitmq.Connection
	logger logger.Logger
	cfg    *config.Config
	ch     *rabbitmq.Channel
}

func New(cfg *config.Config, logger logger.Logger) messager.Messenger {
	return &rabbit{
		logger: logger,
		cfg:    cfg,
	}
}

func (r *rabbit) Write(message string, topic string) error {
	if r.conn == nil || r.conn.IsClosed() {
		r.Connect()
	}

	if r.ch == nil {
		r.Channel()
	}

	if r.ch == nil {
		return errors.New("rabbit connection error")
	}

	err := r.ch.Publish(
		topic, // exchange
		"",    // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         []byte(message),
			DeliveryMode: amqp.Persistent,
		})
	if err != nil {
		return err
	}

	return err
}

func (r *rabbit) Consumer(topic string, callback func(interface{})) error {
	if r.conn == nil || r.conn.IsClosed() {
		r.Connect()
	}

	if r.ch == nil {
		r.Channel()
	}

	if r.ch == nil {
		return errors.New("rabbit connection error")
	}

	msgs, err := r.ch.Consume(
		topic, // queue
		"",    // consumer
		false, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		e := fmt.Sprintf("Failed to register a consumer : %s", err)
		r.logger.Error(errors.New(e))
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			callback(d)
		}
	}()
	<-forever
	return nil
}

func (r *rabbit) Connect() {
	conn, err := rabbitmq.Dial(dns(r.cfg))
	if err != nil {
		r.conn = nil
		r.logger.Error(err)
		return
	}

	r.conn = conn
}

func (r *rabbit) Channel() {
	if r.conn == nil {
		return
	}

	ch, err := r.conn.Channel()
	if err != nil {
		r.logger.Error(err)
		r.ch = nil
		return
	}

	r.ch = ch
}

func dns(c *config.Config) string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s//%s", c.Rabbit.User, c.Rabbit.Password, c.Rabbit.Host, c.Rabbit.Port, c.Rabbit.VHost)
}
