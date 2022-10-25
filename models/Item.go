package models

type Item struct {
	ItemId      int64  `gorm:"primaryKey" json:"item_id"`
	ItemCode    string `gorm:"type:varchar(300)" json:"item_code"`
	Description string `gorm:"type:varchar(300)" json:"description"`
	Quantity    string `gorm:"type:varchar(300)" json:"quantity"`
	OrderId     int64  `gorm:"type:int" json:"order_id"`
}