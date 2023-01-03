package seeds

import (
	"go-rriaudiobook-server/internal/core/entity/models"

	"gorm.io/gorm"
)

func categorySeeder() Seed {
	seeds := []models.Category{
		{Name: "Kebudayaan"},
		{Name: "Entertainment"},
		{Name: "Sport"},
	}
	model := &models.Category{}

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
