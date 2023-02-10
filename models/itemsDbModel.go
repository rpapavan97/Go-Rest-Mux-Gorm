package models

type Item struct {
	ItemID      uint    `json:"item_id" gorm:"primary_key"`
	Description string  `json:"description"`
	Price       float32 `json:"price" gorm:"type:decimal(10,2);"`
	Quantity    int     `json:"quantity"`
	OrderID     uint    `type:"-" `
}
