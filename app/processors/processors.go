package processors

import (
	"context"
	"fmt"
	"horgh-consumer/app/processors/replication"
	"horgh-consumer/app/services"
)

type Processors struct {
	replication Replication
}

type Replication interface {
	Handle(ctx context.Context)
}

func New(database services.Database) Processors {
	fmt.Println("Create processors")
	return Processors{
		replication: replication.New(database),
	}
}
