package kafka

import (
	"context"
	"horgh-consumer/app/config"
)

type Implementation struct {
}

func (i Implementation) Consume(ctx context.Context) error {
	return nil
}

func New(conf config.EventBusConfig) Implementation {
	return Implementation{}
}
