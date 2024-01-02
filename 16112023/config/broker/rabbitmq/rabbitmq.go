package rabbitmq

import (
	"fmt"
	"github.com/kitabisa/kibitalk/config"
	zlog "github.com/rs/zerolog/log"
	"time"

	"context"
	amqp "github.com/rabbitmq/amqp091-go"
)

var RabbitPublish IRabbitPublisher

type IRabbitConsumer interface {
	Consume(queueName string, handler func([]byte)) error
}

type IRabbitPublisher interface {
	Publish(queueName, message string) error
}

// AMQPClient implements the RabbitMQClient interface using the streadway/amqp library.
type AMQPClient struct {
	conn *amqp.Connection
}

func NewAMQPClient() {
	c := config.AppCfg
	fmt.Println(fmt.Sprintf("amqp://%s:%s@%s:%d/", c.RabbitMQ.User, c.RabbitMQ.Pass, c.RabbitMQ.Host, c.RabbitMQ.Port))
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", c.RabbitMQ.User, c.RabbitMQ.Pass, c.RabbitMQ.Host, c.RabbitMQ.Port))
	if err != nil {
		zlog.Fatal().Msgf("Error connecting to RabbitMQ | %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		zlog.Fatal().Msgf("Error creating channel RabbitMQ | %v", err)
	}

	err = ch.Confirm(false)
	if err != nil {
		ch.Close()
		conn.Close()
		zlog.Fatal().Msgf("Error confirming channel RabbitMQ | %v", err)
	}

	RabbitPublish = &AMQPClient{conn: conn}
}

func (c *AMQPClient) Consume(queueName string, handler func([]byte)) error {
	ch, err := c.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			handler(msg.Body)
		}
	}()

	return nil
}

func (c *AMQPClient) Publish(queueName, message string) error {
	ch, err := c.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx, "", queueName, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	})
	if err != nil {
		return err
	}

	return nil
}
