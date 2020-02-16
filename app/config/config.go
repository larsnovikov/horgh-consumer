package config

import (
	"github.com/vrischmann/envconfig"
	"horgh-consumer/app/services/database"
	"horgh-consumer/app/services/database/clickhouse"
	"horgh-consumer/app/services/database/mysql"
	"horgh-consumer/app/services/database/postgresql"
	"horgh-consumer/app/services/database/vertica"
	"horgh-consumer/app/services/eventbus/kafka"
	"horgh-consumer/app/utils/healthcheck"
)

type DatabaseConfig interface{}

type Config struct {
	DatabaseType   string `envconfig:"SLAVE_TYPE"`
	EventBusConfig kafka.Config
	HealthCheck    healthcheck.Config
	DatabaseConfig DatabaseConfig `envconfig:"-"`
}

func New(slaveType string) (Config, error) {
	c := Config{}
	if err := envconfig.Init(&c); err != nil {
		return c, err
	}

	var databaseConfig DatabaseConfig
	switch database.SlaveType(slaveType) {
	case database.TypeClickhouse:
		databaseConfig := clickhouse.Config{}
		if err := envconfig.Init(&databaseConfig); err != nil {
			return c, err
		}
	case database.TypeMysql:
		databaseConfig := mysql.Config{}
		if err := envconfig.Init(&databaseConfig); err != nil {
			return c, err
		}
	case database.TypePostgresql:
		databaseConfig := postgresql.Config{}
		if err := envconfig.Init(&databaseConfig); err != nil {
			return c, err
		}
	case database.TypeVertica:
		databaseConfig := vertica.Config{}
		if err := envconfig.Init(&databaseConfig); err != nil {
			return c, err
		}
	}

	c.DatabaseConfig = databaseConfig

	return c, nil
}
