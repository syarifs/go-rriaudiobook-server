package models

type Book struct {
	ID         uint      `json:"id" gorm:"primary_key"`
	Title      string    `json:"title"`
	Summary    string    `json:"summary"`
	CoverImage string    `json:"cover_image"`
	CategoryID uint      `json:"category_id"`
	Category   Category  `json:"category"`
	Chapter    []Chapter `json:"chapters" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
