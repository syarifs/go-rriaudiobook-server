package models

import (
	"gorm.io/gorm"
)

type Chapter struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Code        uint   `json:"code" gorm:"index"`
	Title       string `json:"title"`
	Description string `json:"description"`
	MediaPath   string `json:"media_path"`
	BookID      uint   `json:"book_id" gorm:"index"`
	Book        Book
}

func (c *Chapter) BeforeCreate(tx *gorm.DB) (err error) {
	var id int64

	tx.Model(&c).
		Where("book_id = ?", c.BookID).
		Count(&id)

	c.Code = uint(id + 1)
	return
}
