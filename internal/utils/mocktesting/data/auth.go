package data

import (
	r "go-rriaudiobook-server/internal/core/entity/request"
	"go-rriaudiobook-server/internal/core/entity/response"
)

var Login r.Login = r.Login{
	Email:    "alsyadahmad@holyhos.co.id",
	Password: "password",
}

var Register r.UserRequest = r.UserRequest{
	ID:       1,
	Email:    "alsyadahmad@holyhos.co.id",
	Password: "password",
	FullName: "Alsyad Ahmad",
	Gender:   "Male",
	RoleID:   2,
}

var ChangePassword r.ChangePassword = r.ChangePassword{
	Code:     "DCR00001",
	Password: "password",
}

var User response.User = response.User{
	ID:       1,
	Code:     "DCR00001",
	Email:    "alsyadahmad@holyhos.co.id",
	Password: "password",
	FullName: "Alsyad Ahmad",
	Role:     "Doctor",
	Status:   1,
}
