package models

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Code     string `json:"code" gorm:"unique"`
	FullName string `json:"full_name"`
	Gender   string `json:"gender"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Status   int    `json:"status" gorm:"default:1"`
	RoleID   int    `json:"role_id"`
	Role     Role   `json:"roles"`

	Book      []Book `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:Code;foreignKey:UserCode;"`
	DeletedAt gorm.DeletedAt
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	var id int64
	var code string

	tx.Model(&u).Unscoped().
		Where("role_id = ?", u.RoleID).
		Count(&id)
	tx.Model(&u.Role).Select("code").
		Where("id = ?", u.RoleID).Scan(&code)
	u.Code = fmt.Sprintf("%s%05d", code, id+1)
	return
}
