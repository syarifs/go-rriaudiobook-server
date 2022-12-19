package repository

import (
	"go-rriaudiobook-server/internal/core/entity/models"
	"go-rriaudiobook-server/internal/core/entity/response"
)

type BookRepository interface {
	FindAll() ([]response.Book, error)
	FindByID(int) (response.BookDetail, error)
	Create(models.Book) error
	InsertChapter(models.Chapter) error
	Update(models.Book) error
	Delete(int) error
}
