package entities

import (
	"time"
)

type Product struct {
	Id       int `gorm:"primaryKey"`
	Name     string
	Price    float64
	Status   bool
	Quantity int

	Created time.Time
}

func (p *Product) TableName() string {
	return "product"
}
