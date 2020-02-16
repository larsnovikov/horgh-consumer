package entities

type HealthCheck struct {
	ServiceName string `json:"service_name"`
	Value       bool   `json:"value"`
}
