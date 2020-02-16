package healthcheck

type Config struct {
	Port int    `envconfig:"HEALTH_CHECK_PORT"`
	Uri  string `envconfig:"HEALTH_CHECK_URI"`
}
