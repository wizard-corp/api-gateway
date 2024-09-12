package myrabbitmq

import (
	"context"
	"errors"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	BROKER_CONNECTION_FAIL            = "failed to connect to Rabbitmq\n"
	BROKER_CONNECTION_CHANNEL_FAIL    = "failed to connect to Channel\n"
	BROKER_DISCONNECTION_FAIL         = "failed to disconnect from Rabbitmq\n"
	BROKER_DISCONNECTION_CHANNEL_FAIL = "failed to disconnect from Channel\n"
)

type RabbitmqConfig struct {
	Host     string
	Port     int
	User     string
	Password string
}

type RabbitmqProducer struct {
	Cli *amqp.Connection
	Cha *amqp.Channel
}

func NewRabbitmqClient(config *RabbitmqConfig) (*RabbitmqProducer, error) {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rabbitmqURI := fmt.Sprintf("amqp://%s:%s@%s:%d/", config.User, config.Password, config.Host, config.Port)
	if config.User == "" || config.Password == "" {
		rabbitmqURI = fmt.Sprintf("mongodb://%s:%d", config.Host, config.Port)
	}

	cli, err := amqp.Dial(rabbitmqURI)
	if err != nil {
		return nil, errors.New(BROKER_CONNECTION_FAIL + err.Error())
	}

	channel, err := cli.Channel()
	if err != nil {
		return nil, errors.New(BROKER_CONNECTION_CHANNEL_FAIL + err.Error())
	}

	return &RabbitmqProducer{Cli: cli, Cha: channel}, nil
}

func (r *RabbitmqProducer) Close() error {
	if r.Cli == nil {
		return nil
	}
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := r.Cha.Close(); err != nil {
		return errors.New(BROKER_DISCONNECTION_FAIL + err.Error())
	}
	if err := r.Cli.Close(); err != nil {
		return errors.New(BROKER_CONNECTION_CHANNEL_FAIL + err.Error())
	}

	return nil
}
