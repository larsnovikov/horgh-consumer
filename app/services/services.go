package services

import (
	"context"
	"horgh-consumer/app/config"
	"horgh-consumer/app/services/database"
	"horgh-consumer/app/services/eventbus"
)

type Implementation struct {
	Database Database
	EventBus EventBus
}

type Database interface {
	Insert(ctx context.Context) error
	Delete(ctx context.Context) error
	Update(ctx context.Context) error
}

type EventBus interface {
	Consume(ctx context.Context) error
}

func New(conf config.Config) Implementation {
	return Implementation{
		Database: database.New(conf.Database),
		EventBus: eventbus.New(conf.EventBusConfig),
	}
}
