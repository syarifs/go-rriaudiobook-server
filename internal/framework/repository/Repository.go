package repository

import (
	"go-rriaudiobook-server/internal/core/repository"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func NewRepository(db *gorm.DB, mongo *mongo.Database) *repository.Repository {
	return &repository.Repository{
		Auth:     NewAuthRepository(db, mongo),
		User:     NewUserRepository(db),
		Book:     NewBookRepository(db),
		Chapter:  NewChapterRepository(db),
		Category: NewCategoryRepository(db),
	}
}
