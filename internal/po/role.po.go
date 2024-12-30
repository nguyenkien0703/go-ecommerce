package po

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID       int64  `gorm:"column:id; type:int;not null;primary_key; required; autoIncrement; comment:'primary Key is ID'"`
	RoleName string `gorm:"column:role_name; type:varchar(255);"`
	//IsActive bool   `gorm:"column:is_active; type:boolean;"`
	RoleNote string `gorm:"column:role_note; type:text;"`
}

func (u *Role) TableName() string {
	return "go_db_role"
}
