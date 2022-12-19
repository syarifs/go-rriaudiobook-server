package response

type (
	User struct {
		ID       uint   `json:"id" example:"1"`
		Code     string `json:"code" example:"DR00001"`
		Email    string `json:"email" example:"alsyadahmad@web.io"`
		FullName string `json:"full_name" example:"Alsyad Ahmad"`
		Password string `json:"-"`
		Role     string `json:"roles" example:"Doctor"`
		Status   int    `json:"status" example:"1"`
	}

	UserDetail struct {
		ID       uint   `json:"id" example:"1"`
		Code     string `json:"code" example:"DR00001"`
		Email    string `json:"email" example:"alsyadahmad@holyhos.co.id"`
		FullName string `json:"full_name" example:"Alsyad Ahmad"`
		Password string `json:"-"`
		Gender   string `json:"gender" example:"Male"`
		Role     string `json:"roles" example:"Doctor"`
		Status   int    `json:"status" example:"1"`
	}
)
