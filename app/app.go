package app

import (
	"horgh-consumer/app/config"
	"horgh-consumer/app/services"
)

type Application struct {
	Services Services
}

type Services interface{}

func New() Application {

	conf := config.Config{}

	return Application{
		Services: services.New(conf),
	}
}
