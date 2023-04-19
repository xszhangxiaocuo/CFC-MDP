// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID       int64  `gorm:"column:id;type:int(64);primaryKey;autoIncrement:true" json:"id,string"`
	Account  string `gorm:"column:account;type:varchar(255);not null" json:"account"`
	Password string `gorm:"column:password;type:varchar(255);not null" json:"password"`
	Deleted  int64  `gorm:"column:deleted;type:int(1) unsigned zerofill" json:"deleted"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
