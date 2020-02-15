package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"horgh-consumer/app/entities"
	"horgh-consumer/app/utils/logger"
)

type Implementation struct {
	readers []*kafka.Reader
}

func (i Implementation) Consume(ctx context.Context, handler func(ctx context.Context, message entities.Query) error) error {
	for _, reader := range i.readers {
		go i.consume(ctx, reader, handler)
	}

	return nil
}

func (i Implementation) consume(ctx context.Context, reader *kafka.Reader, handler func(ctx context.Context, message entities.Query) error) {
	for {
		i.handle(ctx, reader, handler)
	}
}

func (i Implementation) handle(ctx context.Context, reader *kafka.Reader, handler func(ctx context.Context, message entities.Query) error) {
	simpleLogger := logger.Get(ctx)
	var outErr error

	defer func() {
		if r := recover(); r != nil {
			outErr = r.(error)
		}

		if outErr != nil {
			simpleLogger.Error(fmt.Sprintf("Handle message error: %s", outErr.Error()))
		}
	}()

	m, err := reader.FetchMessage(context.Background())
	if err != nil {
		outErr = err
		return
	}

	msg, err := entities.Parse(string(m.Value))
	if err != nil {
		outErr = err
		return
	}

	if err = handler(ctx, msg); err != nil {
		outErr = err
		return
	}

	outErr = reader.CommitMessages(ctx, m)
}

func New(conf Config) Implementation {
	var readers []*kafka.Reader
	for _, topic := range conf.Topics {
		fmt.Println("Create reader for " + topic)
		readers = append(readers, kafka.NewReader(kafka.ReaderConfig{
			Brokers: conf.Hosts,
			GroupID: conf.ConsumerGroup,
			Topic:   topic,
		}))
	}

	return Implementation{
		readers: readers,
	}
}
