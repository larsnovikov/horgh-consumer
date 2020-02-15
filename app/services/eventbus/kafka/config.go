package kafka

type Config struct {
	Hosts         []string `envconfig:"Hosts"`
	Topics        []string `envconfig:"Topics"`
	ConsumerGroup string   `envconfig:"ConsumerGroup"`
}
