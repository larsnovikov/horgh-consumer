package app

import (
	"context"
	"fmt"
	"github.com/vrischmann/envconfig"
	"horgh-consumer/app/config"
	"horgh-consumer/app/processors"
	"horgh-consumer/app/services/database"
	"horgh-consumer/app/services/eventbus"
)

type Application struct {
	Transport Transport
}

type Transport interface{}

func New() (Application, error) {
	ctx := context.Background()

	conf := config.Config{}
	if err := envconfig.Init(&conf); err != nil {
		return Application{}, err
	}

	var tmpDbConfig struct{}
	db := database.New(tmpDbConfig)
	proc := processors.New(db)
	eb := eventbus.New(conf.EventBusConfig)

	if err := eb.Consume(ctx, proc.Replication.Handle); err != nil {
		return Application{}, err
	}

	return Application{
		Transport: eb,
	}, nil
}

func (app Application) Wait() {
	var a string
	fmt.Scanln(&a)
}
