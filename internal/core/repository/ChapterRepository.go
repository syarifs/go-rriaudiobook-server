package repository

import "go-rriaudiobook-server/internal/core/entity/models"

type ChapterRepository interface {
	Create(models.Chapter) error
	Update(models.Chapter) error
	Delete(int, int) error
	GetAudio(int, int) (string, error)
}
