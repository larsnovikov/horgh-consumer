package services

import (
	"context"
)

type Database interface {
	Insert(ctx context.Context) error
	Delete(ctx context.Context) error
	Update(ctx context.Context) error
}

type EventBus interface {
	Consume(ctx context.Context) error
}
