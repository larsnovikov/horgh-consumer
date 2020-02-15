package replication

import (
	"context"
	"horgh-consumer/app/entities"
	"horgh-consumer/app/services"
)

type Implementation struct {
	database services.Database
}

func (i Implementation) Handle(ctx context.Context, message entities.Query) error {
	return nil
}

func New(database services.Database) Implementation {
	return Implementation{
		database: database,
	}
}
