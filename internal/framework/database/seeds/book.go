package seeds

import (
	"go-rriaudiobook-server/internal/core/entity/models"

	"gorm.io/gorm"
)

func bookSeeder() Seed {
	seeds := []models.Book{
		{
			Title:   "Test 1",
			Summary: "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.",
			Category: models.Category{
				Name: "Test Category",
			},
		},
		{
			Title:   "Test 2",
			Summary: "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.",
			Category: models.Category{
				Name: "Test Category",
			},
		},
	}
	model := &models.Book{}

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
