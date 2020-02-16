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

func (i Implementation) HealthCheck() entities.HealthCheck {
	return i.client.HealthCheck()
}

type Client interface {
	Consume(ctx context.Context, handler func(ctx context.Context, message entities.Query) error) error
	HealthCheck() entities.HealthCheck
}

func New(conf kafka.Config) (Implementation, error) {
	kafkaClient, err := kafka.New(conf)
	if err != nil {
		return Implementation{}, err
	}
	return Implementation{
		client: kafkaClient,
	}, nil
}
