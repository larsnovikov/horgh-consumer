package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
)

type Implementation struct {
	readers []*kafka.Reader
}

func (i Implementation) Consume(ctx context.Context) error {
	for _, reader := range i.readers {
		go i.consume(ctx, reader)
	}

	return nil
}

func (i Implementation) consume(ctx context.Context, reader *kafka.Reader) {
	for {
		i.handle(ctx, reader)
	}
}

func (i Implementation) handle(ctx context.Context, reader *kafka.Reader) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	m, err := reader.ReadMessage(context.Background())
	if err != nil {
		// todo log
		return
	}

	fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
}

func New(conf Config) Implementation {
	var readers []*kafka.Reader
	for _, topic := range conf.Topics {
		fmt.Println("Create reader for " + topic)
		readers = append(readers, kafka.NewReader(kafka.ReaderConfig{
			Brokers:  conf.Hosts,
			GroupID:  conf.ConsumerGroup,
			Topic:    topic,
			MinBytes: 10e3, // 10KB
			MaxBytes: 10e6, // 10MB
		}))
	}

	return Implementation{
		readers: readers,
	}
}
