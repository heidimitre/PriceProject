package price

import "time"

type Price struct {
	Value float64
	TimeRecorded time.Time
}

func CreatePrice (newValue float64, currentTime time.Time) (price *Price){
	newPrice := new(Price)
	newPrice.Value = newValue
	newPrice.TimeRecorded = currentTime
	return newPrice
}