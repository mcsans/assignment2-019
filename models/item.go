package models

import "time"

type Item struct {
	ItemId      uint			`gorm:"primaryKey"`
	ItemCode    string		`gorm:"type:varchar(100)" json:"itemCode" binding:"required"`
	Description string		`gorm:"type:text" json:"description" binding:"required"`
	Quantity    int				`json:"quantity" binding:"required"`
	OrderId     uint			
	CreatedAt   time.Time
	UpdatedAt   time.Time
}