package eventbus

import (
	"context"
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

func New(conf kafka.Config) Implementation {
	return Implementation{
		client: kafka.New(conf),
	}
}
