package seeds

import (
	"go-rriaudiobook-server/internal/core/entity/models"
	"go-rriaudiobook-server/internal/utils"

	"gorm.io/gorm"
)

func userSeeder() Seed {
	password, _ := utils.HashPassword("password")
	seeds := []models.User{
		{
			FullName: "Rimuru Tempest",
			Email:    "rimurutempest@system.co.id",
			Gender:   "Male",
			Password: password,
			RoleID:   1,
			Status:   1,
		},
		{
			FullName: "Alsyad Ahmad",
			Email:    "alsyadahmad@system.co.id",
			Password: password,
			Gender:   "Male",
			RoleID:   2,
			Status:   1,
		},
		{
			FullName: "Milim Nava",
			Email:    "milimnava@system.co.id",
			Password: password,
			Gender:   "Female",
			RoleID:   2,
			Status:   1,
		},
	}
	model := &models.User{}

	return Seed{
		models: model,
		run: func(db *gorm.DB) (err error) {
			for _, v := range seeds {
				err = db.Create(&v).Error
			}
			return
		},
	}
}
