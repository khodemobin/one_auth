package broker

import (
	"errors"
	"fmt"
	"github.com/isayme/go-amqp-reconnect/rabbitmq"
	"github.com/khodemobin/pilo/auth/config"
	"github.com/streadway/amqp"
)

type Broker interface {
	Write(message, topic string) error
	Consumer(topic string, callback func(msg interface{})) error
}

type rabbit struct {
	conn *rabbitmq.Connection
	cfg  *config.Config
	ch   *rabbitmq.Channel
}

func NewRabbitMQ(cfg *config.Config) Broker {
	return &rabbit{
		cfg: cfg,
	}
}

func (r *rabbit) Write(message string, topic string) error {
	if r.conn == nil || r.conn.IsClosed() {
		if err := r.Connect(); err != nil {
			return err
		}
	}

	if r.ch == nil {
		if err := r.Channel(); err != nil {
			return err
		}
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
		if err := r.Connect(); err != nil {
			return err
		}
	}

	if r.ch == nil {
		if err := r.Channel(); err != nil {
			return err
		}
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
		return errors.New(e)
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

func (r *rabbit) Connect() error {
	conn, err := rabbitmq.Dial(dns(r.cfg))
	if err != nil {
		r.conn = nil
		return err
	}

	r.conn = conn

	return nil
}

func (r *rabbit) Channel() error {
	if r.conn == nil {
		return nil
	}

	ch, err := r.conn.Channel()
	if err != nil {
		r.ch = nil
		return err
	}

	r.ch = ch

	return nil
}

func dns(c *config.Config) string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s//%s", c.Rabbit.User, c.Rabbit.Password, c.Rabbit.Host, c.Rabbit.Port, c.Rabbit.VHost)
}
