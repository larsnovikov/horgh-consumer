package config

import "horgh-consumer/app/services/eventbus/kafka"

type DatabaseConfig interface{}

type Config struct {
	//Database       DatabaseConfig
	EventBusConfig kafka.Config
}
