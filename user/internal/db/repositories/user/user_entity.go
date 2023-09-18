package user

import (
	"time"
)

type User struct {
	ID        string    `gorm:"column:id;type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name      string    `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Email     string    `gorm:"column:email;type:varchar(255);not null" json:"email"`
	Password  string    `gorm:"column:password;type:varchar(255);not null" json:"password"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// set table name
func (User) TableName() string {
	return "user"
}
