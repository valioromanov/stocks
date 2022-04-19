package models

type Symbol struct {
	SymbolId string `db:"symbol"`
	Status   string
}

type SymbolInterface interface {
	FindAll() ([]Symbol, error)
	FindAllByStatis() ([]Symbol, error)
}
