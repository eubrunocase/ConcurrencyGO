package models

import "time"

type PriceDetails struct {
	StoreName string
	Value     float64
	TimeStamp time.Time
}