package vertica

import (
	"context"
)

type Implementation struct {
}

func (i Implementation) Insert(ctx context.Context) error {
	return nil
}

func (i Implementation) Delete(ctx context.Context) error {
	return nil
}

func (i Implementation) Update(ctx context.Context) error {
	return nil
}

func New(conf Config) Implementation {
	return Implementation{}
}
