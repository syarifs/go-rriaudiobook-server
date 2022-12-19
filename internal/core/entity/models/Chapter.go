package models

type Chapter struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Title     string `json:"title"`
	MediaPath string `json:"media_path"`
	BookID    uint   `json:"book_id"`
}
