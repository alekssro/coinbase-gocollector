package domain

import "time"

type HistoricRate struct {
	Time   time.Time `db:"date"`
	Low    float64   `db:"low"`
	High   float64   `db:"high"`
	Open   float64   `db:"open"`
	Close  float64   `db:"close"`
	Volume float64   `db:"volume"`
}
