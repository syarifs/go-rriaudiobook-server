package request

import (
	vl "go-rriaudiobook-server/internal/utils/validators"

	v "github.com/go-ozzo/ozzo-validation/v4"
)

type BookRequest struct {
	ID         uint   `json:"id" form:"id" example:"1"`
	Title      string `json:"title" form:"title" example:"book 1"`
	Summary    string `json:"summary" form:"summary" example:"Lorem ipsum dolor sit amet."`
	CoverImage string `json:"cover_image" form:"cover_image" example:""`
	CategoryID uint   `json:"category_id" form:"category_id" example:"1"`
}

func (br BookRequest) Validate() (err error) {
	err = v.ValidateStruct(&br,
		v.Field(&br.Title, v.Required, v.By(vl.Duplicate("books", "title", int(br.ID)))),
		v.Field(&br.Summary, v.Required),
		v.Field(&br.CoverImage, v.Required),
		v.Field(&br.CategoryID, v.Required),
	)
	return
}
