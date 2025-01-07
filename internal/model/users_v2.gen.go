package model

import (
	"time"
)

const TableNameUsersV2 = "users_v2"

// UsersV2 mapped from table <users_v2>
type UsersV2 struct {
	UserID    int32     `gorm:"column:user_id;primaryKey;autoIncrement:true" json:"user_id"`
	NameV2    string    `gorm:"column:name;not null" json:"name_v2"`
	Email     string    `gorm:"column:email;not null" json:"email"`
	Password  string    `gorm:"column:password;not null" json:"password"`
	Phone     string    `gorm:"column:phone" json:"phone"`
	Role      string    `gorm:"column:role;default:user" json:"role"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
}

// TableName UsersV2's table name
func (*UsersV2) TableName() string {
	return TableNameUsersV2
}
