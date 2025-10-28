package domain

import "time"

type Product struct {
	ID           int
	ProductName  string
	Manufacturer string
	ProductCount int
	Price        float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
