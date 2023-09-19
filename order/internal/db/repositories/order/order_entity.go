package order

import (
	"time"
)

type Order struct {
	ID            int       `gorm:"column:id;type:uuid;primary_key" json:"id"`
	ParentOrderID *int      `gorm:"column:parent_order_id;type:uuid" json:"parent_order_id"`
	UserID        string    `gorm:"column:user_id;type:uuid" json:"user_id"`
	ProductId     string    `gorm:"column:product_id;type:uuid" json:"product_id"`
	Quantity      int64     `gorm:"column:quantity" json:"quantity"`
	Price         int64     `gorm:"column:price" json:"price"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// set table name
func (Order) TableName() string {
	return "order"
}
