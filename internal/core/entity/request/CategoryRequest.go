package request

import (
	vl "go-rriaudiobook-server/internal/utils/validators"

	v "github.com/go-ozzo/ozzo-validation/v4"
)

type CategoryRequest struct {
	ID   uint   `json:"id" example:"1"`
	Name string `json:"name" example:"Kebudayaan"`
}

func (cr CategoryRequest) Validate() (err error) {
	err = v.ValidateStruct(&cr,
		v.Field(&cr.Name, v.Required, v.By(vl.Duplicate("categories", "name", int(cr.ID)))),
	)
	return
}
