package models

type Record struct {
	Query string `db:"query"`
	Freq  int    `db:"freq"`
}
