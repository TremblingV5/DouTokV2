// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID              int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	AccountID       int64     `gorm:"column:account_id" json:"account_id"`
	Mobile          string    `gorm:"column:mobile" json:"mobile"`
	Email           string    `gorm:"column:email" json:"email"`
	Name            string    `gorm:"column:name" json:"name"`
	Avatar          string    `gorm:"column:avatar" json:"avatar"`
	BackgroundImage string    `gorm:"column:background_image" json:"background_image"`
	Signature       string    `gorm:"column:signature" json:"signature"`
	CreatedAt       time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
