package main

import (
	"encoding/csv"
	"fmt"
	"github.com/ereminiu/autocomplete/models"
	"github.com/ereminiu/autocomplete/pkg/repository"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

const filepath = "demo-data/large.csv"
const smallpath = "demo-data/records.csv"

func FetchRecords() []models.Record {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	reader := csv.NewReader(f)
	records := make([]models.Record, 0)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		freq, err := strconv.Atoi(record[1])
		if err != nil {
			log.Fatal(err)
		}
		records = append(records, models.Record{Query: record[0], Freq: freq})
	}
	return records
}

func AddAllRecords(rep *repository.Repository) error {
	records := FetchRecords()
	return rep.BatchInsert(records)
}

func main() {
	start := time.Now()

	rep, err := repository.NewRepository()
	if err != nil {
		log.Fatal(err)
	}
	if err := AddAllRecords(rep); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("time elapsed: %v\n", time.Since(start))
}
