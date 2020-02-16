package config

import (
	"horgh-consumer/app/services/eventbus/kafka"
	"horgh-consumer/app/utils/healthcheck"
)

type DatabaseConfig interface{}

type Config struct {
	//Database       DatabaseConfig
	EventBusConfig kafka.Config
	HealthCheck    healthcheck.Config
}
