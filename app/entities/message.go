package entities

type Column struct {
	Name  string
	Value interface{}
}

type Query struct {
	columns []Column
}
