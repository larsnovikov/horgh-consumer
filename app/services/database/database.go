package database

import (
	"context"
)

const (
	TypeMysql      = SlaveType("mysql")
	TypePostgresql = SlaveType("postgresql")
	TypeVertica    = SlaveType("vertica")
	TypeClickhouse = SlaveType("clickhouse")
)

type SlaveType string

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

func New(conf interface{}) Implementation {
	i := Implementation{}
	// TODO
	return i
}
