package eventbus

import (
	"context"
	"horgh-consumer/app/entities"
	"horgh-consumer/app/services/eventbus/kafka"
)

type Implementation struct {
	client Client
}

func (i Implementation) Consume(ctx context.Context, handler func(ctx context.Context, message entities.Query) error) error {
	return i.client.Consume(ctx, handler)
}

type Client interface {
	Consume(ctx context.Context, handler func(ctx context.Context, message entities.Query) error) error
}

func New(conf kafka.Config) Implementation {
	return Implementation{
		client: kafka.New(conf),
	}
}
