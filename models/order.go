package models

import "time"

type Order struct {
	OrderId      uint				`gorm:"primaryKey"`
	CustomerName string			`gorm:"type:varchar(100)" json:"customerName" binding:"required"`
	OrderedAt    time.Time	`json:"orderedAt" binding:"required"`
	CreatedAt    time.Time	
	UpdatedAt    time.Time	
	Items        []Item			`json:"items" binding:"required"`
}