package vertica

type Config struct {
	Host     string `envconfig:"SLAVE_HOST"`
	Port     int    `envconfig:"SLAVE_PORT"`
	User     string `envconfig:"SLAVE_USER"`
	Password string `envconfig:"SLAVE_PASSWORD"`
}
