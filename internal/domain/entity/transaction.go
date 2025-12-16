package entity

import "time"

type Transaction struct {
	PAN        string
	ExpiryDate time.Time
	CVM        string
	Status     string
	CreatedAt  time.Time
}
