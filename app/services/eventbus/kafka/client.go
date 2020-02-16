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
	conn    *kafka.Conn
}

func (i Implementation) HealthCheck() entities.HealthCheck {
	brokers, err := i.conn.Brokers()

	return entities.HealthCheck{
		ServiceName: "kafka",
		Value:       len(brokers) > 0 && err == nil,
	}
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

func New(conf Config) (Implementation, error) {
	var readers []*kafka.Reader
	var conn *kafka.Conn
	var err error

	for _, topic := range conf.Topics {
		fmt.Println("Create reader for " + topic)
		readers = append(readers, kafka.NewReader(kafka.ReaderConfig{
			Brokers: conf.Hosts,
			GroupID: conf.ConsumerGroup,
			Topic:   topic,
		}))

		if conn == nil {
			conn, err = kafka.DialLeader(context.Background(), "tcp", conf.Hosts[0], topic, 0)
			if err != nil {
				return Implementation{}, err
			}
		}
	}

	return Implementation{
		readers: readers,
		conn:    conn,
	}, nil
}
