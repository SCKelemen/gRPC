package Corpus

type Corpus struct {
}

type CorpusType int

const (
	Code CorpusType = iota
	Employees
	Teams
	Products
)
