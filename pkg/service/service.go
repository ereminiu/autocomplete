package service

import (
	"github.com/ereminiu/autocomplete/models"
	"github.com/ereminiu/autocomplete/pkg/repository"
	"github.com/ereminiu/autocomplete/trie"
)

type Autocomplete interface {
	AddWord(word string)
	GetTopFive(prefix string) []string
	Rebuild(records []models.Record)
}

type Repository interface {
	AddRecord(rec models.Record) error
	FetchRecords() ([]models.Record, error)
	BatchInsert(records []models.Record) error
}

type Service struct {
	Autocomplete
	Repository
}

func NewService(rep *repository.Repository, trie *trie.Trie) *Service {
	return &Service{Repository: rep, Autocomplete: trie}
}
