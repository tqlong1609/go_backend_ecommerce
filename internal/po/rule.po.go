package po

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID   int64  `gorm:"column:uuid; type:int; not null; index:idx_uuid; primaryKey; autooIncrement; comment:'primary Key is primary key'"`
	Name string `gorm:"column: name; type:varchar(255); not null;"`
	Note string `gorm:"column: note; type:text; "`
}

func (r *Role) TableName() string {
	return "roles"
}
