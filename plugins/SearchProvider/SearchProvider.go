package SearchProvider

import (
	"github.com/SCKelemen/grpc/Corpus"
)

type ISearchProvider interface {
	Query(term string) ([]string, error)
	GetCorpora() ([]string, error)
	SearchesCorpus(corpora ...string) bool
	SearchesCorpora(corpora []string) bool
}

type SearchProvider struct {
	Corpora map[string]Corpus.Corpus
}

func (sp SearchProvider) Query(term string) ([]string, error) {
	return nil, nil
}

func (sp SearchProvider) GetCorpora() ([]string, error) {
	return nil, nil
}

func (sp SearchProvider) SearchesCorpus(corpora ...string) bool {
	for corpus := range sp.Corpora {
		_, ok := sp.Corpora[corpus]
		if !ok {
			return false
		}
	}
	return true
}

type SymbolCorpus struct {
	Indexer  string
	Searcher ISearchProvider
}
