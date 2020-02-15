package entities

import "encoding/json"

type Column struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

type Query struct {
	Data []Column `json:"data"`
}

func Parse(message string) (Query, error) {
	var query Query
	err := json.Unmarshal([]byte(message), &query)

	return query, err
}
