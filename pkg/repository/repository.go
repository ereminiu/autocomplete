package repository

import (
	"fmt"
	"github.com/ereminiu/autocomplete/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"strings"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository() (*Repository, error) {
	dataSource := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
		"ys-user", "qwerty", "localhost", 5432, "ys-db",
	)
	conn, err := sqlx.Connect("postgres", dataSource)
	if err != nil {
		return nil, errors.Wrap(err, "sqlx connect")
	}
	if err := conn.Ping(); err != nil {
		return nil, errors.Wrap(err, "ping failed")
	}
	return &Repository{db: conn}, nil
}

func (r *Repository) AddRecord(rec models.Record) error {
	query := `INSERT INTO ftable (query, freq) 
			  VALUES ($1, $2) 
			  ON CONFLICT (query)
			  DO UPDATE SET freq = ftable.freq+1
			  RETURNING ID`
	_, err := r.db.Exec(query, rec.Query, rec.Freq)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) FetchRecords() ([]models.Record, error) {
	query := `SELECT query, freq FROM ftable`
	var records []models.Record
	err := r.db.Select(&records, query)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (r *Repository) addRecords(records []models.Record) error {
	valuesQuery := make([]string, 0, len(records))
	valuesArgs := make([]interface{}, 0, len(records))
	i := 1
	for _, r := range records {
		valuesQuery = append(valuesQuery, fmt.Sprintf("($%d, $%d)", i, i+1))
		valuesArgs = append(valuesArgs, r.Query, r.Freq)
		i += 2
	}
	query := fmt.Sprintf(`
		INSERT INTO ftable (query, freq)
		VALUES %s 
		ON CONFLICT(query) DO
		UPDATE SET freq = ftable.freq+1`, strings.Join(valuesQuery, ", "))
	_, err := r.db.Exec(query, valuesArgs...)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) BatchInsert(records []models.Record) error {
	batch := 250
	n := len(records)
	for i := 0; i < n; i += batch {
		last := min(i+batch, n)
		err := r.addRecords(records[i:last])
		if err != nil {
			return errors.Wrap(err, "batch insert")
		}
	}
	return nil
}
