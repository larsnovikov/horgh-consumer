package eventbus

import (
	"context"
	"horgh-consumer/app/config"
	"horgh-consumer/app/services/eventbus/kafka"
)

type Implementation struct {
	client Client
}

func (i Implementation) Consume(ctx context.Context) error {
	return i.client.Consume(ctx)
}

type Client interface {
	Consume(ctx context.Context) error
}

func New(conf config.EventBusConfig) Implementation {
	return Implementation{
		client: kafka.New(conf),
	}
}
