package healthcheck

import (
	"encoding/json"
	"fmt"
	"horgh-consumer/app/entities"
	"net/http"
)

type Implementation struct {
	uri     string
	port    int
	methods []func() entities.HealthCheck
}

type Output struct {
	Services []entities.HealthCheck `json:"services"`
}

func (o *Output) String() (string, error) {
	b, err := json.Marshal(o)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func New(conf Config, methods []func() entities.HealthCheck) Implementation {
	return Implementation{
		uri:     conf.Uri,
		port:    conf.Port,
		methods: methods,
	}
}

func (i Implementation) Handle() {
	http.HandleFunc(i.uri, func(w http.ResponseWriter, r *http.Request) {
		output := Output{}
		for _, method := range i.methods {
			res := method()
			output.Services = append(output.Services, res)
		}

		out, err := output.String()
		if err != nil {
			// TODO log
		}

		fmt.Fprintf(w, out)
	})

	http.ListenAndServe(fmt.Sprintf(":%d", i.port), nil)
}
