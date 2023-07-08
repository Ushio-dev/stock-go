package models

type Item struct {
	ItemID uint `gorm:"primaryKey"`
	Name   string
	Amount uint
}
