package replication

import (
	"context"
	"horgh-consumer/app/services"
)

type Implementation struct {
	database services.Database
}

func (i Implementation) Handle(ctx context.Context) {

}

func New(database services.Database) Implementation {
	return Implementation{
		database: database,
	}
}
