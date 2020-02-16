package kafka

type Config struct {
	Hosts         []string `envconfig:"KAFKA_HOSTS"`
	Topics        []string `envconfig:"KAFKA_TOPICS"`
	ConsumerGroup string   `envconfig:"KAFKA_CONSUMER_GROUP"`
}
