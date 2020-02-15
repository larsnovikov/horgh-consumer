package app

import (
	"fmt"
	"github.com/vrischmann/envconfig"
	"horgh-consumer/app/config"
	"horgh-consumer/app/processors"
	"horgh-consumer/app/services/database"
	"horgh-consumer/app/services/eventbus"
)

type Application struct {
	Processors Processors
	Transport  Transport
}

type Processors interface{}
type Transport interface{}

func New() (Application, error) {
	conf := config.Config{}
	if err := envconfig.Init(&conf); err != nil {
		return Application{}, err
	}

	var tmpDbConfig struct{}
	db := database.New(tmpDbConfig)

	return Application{
		Transport:  eventbus.New(conf.EventBusConfig),
		Processors: processors.New(db),
	}, nil
}

func (app Application) Wait() {
	var a string
	fmt.Scanln(&a)
}
