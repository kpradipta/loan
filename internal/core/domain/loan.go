package domain

import (
	"time"
)

type Loan struct{
	Id         string `json:"id"`
	Amount     int64 `json:"amount"`
	BorrowerId string `json: "borrower_id"`
	Tenor      int16  `json: "tenor"`
	CreatedAt  time.Time `json: "created_at"`
	LastUpdate time.Time `json: "last_update"`
	UpdateBy   time.Time `json: "update_by"`
	State      string
}