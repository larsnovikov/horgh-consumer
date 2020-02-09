package config

type DatabaseConfig interface {
	Type() string
	Host() string
	User() string
	Password() string
	Port() int
}

type EventBusConfig interface{}

type Config struct {
	Database       DatabaseConfig
	EventBusConfig EventBusConfig
}
