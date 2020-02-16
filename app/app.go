package app

import (
	"context"
	"github.com/vrischmann/envconfig"
	"horgh-consumer/app/config"
	"horgh-consumer/app/entities"
	"horgh-consumer/app/processors"
	"horgh-consumer/app/services/database"
	"horgh-consumer/app/services/eventbus"
	"horgh-consumer/app/utils/healthcheck"
	"horgh-consumer/app/utils/logger"
)

type Application struct {
	transport   Transport
	healthCheck HealthCheck
}

type Transport interface{}
type HealthCheck interface {
	Handle()
}

func New() (Application, error) {
	ctx := context.Background()
	ctx, err := logger.Set(ctx)
	if err != nil {
		return Application{}, err
	}

	conf := config.Config{}
	if err := envconfig.Init(&conf); err != nil {
		return Application{}, err
	}

	var tmpDbConfig struct{}
	db := database.New(tmpDbConfig)
	proc := processors.New(db)
	eb, err := eventbus.New(conf.EventBusConfig)
	if err != nil {
		return Application{}, err
	}

	if err := eb.Consume(ctx, proc.Replication.Handle); err != nil {
		return Application{}, err
	}

	return Application{
		transport: eb,
		healthCheck: healthcheck.New(conf.HealthCheck, []func() entities.HealthCheck{
			eb.HealthCheck,
		}),
	}, nil
}

func (app Application) HealthCheck() {
	app.healthCheck.Handle()
}
