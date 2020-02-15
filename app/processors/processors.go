package processors

import (
	"context"
	"fmt"
	"horgh-consumer/app/entities"
	"horgh-consumer/app/processors/replication"
	"horgh-consumer/app/services"
)

type Processors struct {
	Replication Replication
}

type Replication interface {
	Handle(ctx context.Context, message entities.Query) error
}

func New(database services.Database) Processors {
	fmt.Println("Create processors")
	return Processors{
		Replication: replication.New(database),
	}
}
