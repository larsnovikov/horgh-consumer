package database

import (
	"context"
	"horgh-consumer/app/config"
)

const (
	typeMysql      = slaveType("mysql")
	typePostgresql = slaveType("postgresql")
	typeVertica    = slaveType("vertica")
	typeClickhouse = slaveType("clickhouse")
)

type slaveType string

type Implementation struct {
	client Client
}

func (i Implementation) Insert(ctx context.Context) error {
	return i.client.Insert(ctx)
}

func (i Implementation) Delete(ctx context.Context) error {
	return i.client.Delete(ctx)
}

func (i Implementation) Update(ctx context.Context) error {
	return i.client.Update(ctx)
}

type Client interface {
	Insert(ctx context.Context) error
	Delete(ctx context.Context) error
	Update(ctx context.Context) error
}

func New(conf config.DatabaseConfig) Implementation {
	i := Implementation{}
	// TODO
	return i
	//switch slaveType(conf.Type()) {
	//case typeClickhouse:
	//	i.client = clickhouse.New(conf)
	//case typeVertica:
	//	i.client = vertica.New(conf)
	//case typeMysql:
	//	i.client = mysql.New(conf)
	//case typePostgresql:
	//	i.client = postgresql.New(conf)
	//}
	//
	//return i
}
