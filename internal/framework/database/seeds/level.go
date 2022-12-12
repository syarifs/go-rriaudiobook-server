package seeds

import (
	"go-rriaudiobook-server/internal/core/entity/models"

	"gorm.io/gorm"
)

func levelSeeder() Seed {
	seeds := []models.Role{
		{Name: "Admin", Code: "ADM"},
		{Name: "Uploader", Code: "UPD"},
		{Name: "Regular", Code: "REG"},
	}
	model := &models.Role{}

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
