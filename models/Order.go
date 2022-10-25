package models

type Order struct {
	OrderId      int64  `gorm:"primaryKey" json:"order_id"`
	CustomerName string `gorm:"type:varchar(300)" json:"customer_name"`
	OrderedAt    string `gorm:"type:varchar(300)" json:"ordered_at"`
}