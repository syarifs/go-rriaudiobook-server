package request

import (
	vl "go-rriaudiobook-server/internal/utils/validators"

	v "github.com/go-ozzo/ozzo-validation/v4"
)

type ChapterRequest struct {
	ID          uint   `json:"id" form:"id" example:"1"`
	Title       string `json:"title" form:"title" example:"chapter 1"`
	Description string `json:"description" form:"description" example:"chapter 1 description"`
	MediaPath   string `json:"media_path" form:"media_path" example:""`
	BookID      uint   `json:"book_id" form:"book_id" example:"2"`
}

func (cr ChapterRequest) Validate() (err error) {
	err = v.ValidateStruct(&cr,
		v.Field(&cr.Title, v.Required, v.By(vl.Duplicate("chapters", "title", int(cr.ID)))),
		v.Field(&cr.Description, v.Required),
		v.Field(&cr.MediaPath, v.Required),
		v.Field(&cr.BookID, v.Required),
	)
	return
}
