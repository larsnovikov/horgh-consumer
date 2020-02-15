package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"horgh-consumer/app/entities"
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
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	m, err := reader.ReadMessage(context.Background())
	if err != nil {
		// todo log
		fmt.Println("222")
		return
	}

	msg, err := entities.Parse(string(m.Value))
	if err != nil {
		fmt.Println(err)
		fmt.Println(string(m.Value))
		// todo log
		return
	}

	if err = handler(ctx, msg); err != nil {
		// todo log
		fmt.Println("222")
		return
	}

	fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
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
