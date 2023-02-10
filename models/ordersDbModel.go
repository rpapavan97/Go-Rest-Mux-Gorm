package models

import "time"

type Order struct {
	OrderID      uint       `json:"order_id" gorm:"primary_key"`
	OrderedAt    time.Time  `json:"ordered_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt    *time.Time `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
	Status       string     `json:"status"`
	Total        float32    `json:"total" gorm:"type:decimal(10,2);"`
	CurrencyUnit string     `json:"currency_unit"`
	Items        []Item     `json:"items" gorm:"foreignkey:OrderID"`
}
