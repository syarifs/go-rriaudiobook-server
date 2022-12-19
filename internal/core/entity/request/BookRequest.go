package request

import (
	v "github.com/go-ozzo/ozzo-validation/v4"
)

type BookRequest struct {
	ID         uint   `json:"id" form:"id" example:"1"`
	Title      string `json:"title" form:"title" example:"book 1"`
	Summary    string `json:"summary" form:"summary" example:"Lorem ipsum dolor sit amet."`
	CoverImage string `json:"cover_image" form:"cover_image" example:""`
	CategoryID uint   `json:"category_id" form:"category_id" example:"1"`
}

type ChapterRequest struct {
	ID        uint   `json:"id" form:"id" example:"1"`
	Title     string `json:"title" form:"title" example:"chapter 1"`
	MediaPath string `json:"media_path" form:"media_path" example:""`
	BookID    uint   `json:"book_id" form:"book_id" example:"2"`
}

func (br BookRequest) Validate() (err error) {
	err = v.ValidateStruct(&br,
		v.Field(&br.Title, v.Required),
		v.Field(&br.Summary, v.Required),
		v.Field(&br.CoverImage, v.Required),
		v.Field(&br.CategoryID, v.Required),
	)
	return
}

func (cr ChapterRequest) Validate() (err error) {
	err = v.ValidateStruct(&cr,
		v.Field(&cr.Title, v.Required),
		v.Field(&cr.MediaPath, v.Required),
		v.Field(&cr.BookID, v.Required),
	)
	return
}
