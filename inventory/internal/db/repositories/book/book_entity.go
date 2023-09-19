package book

import (
	"time"
)

type Book struct {
	ID        string    `gorm:"column:id;type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Title     string    `gorm:"column:title;type:varchar(255)" json:"title"`
	Author    string    `gorm:"column:author;type:varchar(255)" json:"author"`
	Price     int64     `gorm:"column:price;type:bigint" json:"price"` // price in cents (USD)
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// set table name
func (Book) TableName() string {
	return "book"
}
