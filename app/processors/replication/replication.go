package replication

import (
	"context"
	"fmt"
	"horgh-consumer/app/entities"
	"horgh-consumer/app/services"
)

type Implementation struct {
	database services.Database
}

func (i Implementation) Handle(ctx context.Context, message entities.Query) error {
	fmt.Println(message.Data[0].Name)
	fmt.Println(message.Data[0].Value)
	return nil
}

func New(database services.Database) Implementation {
	return Implementation{
		database: database,
	}
}
