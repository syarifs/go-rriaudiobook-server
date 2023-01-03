package repository

import (
	"go-rriaudiobook-server/internal/core/entity/models"
	"go-rriaudiobook-server/internal/core/entity/response"
)

type BookRepository interface {
	FindAll() ([]response.Book, error)
	FindByID(int) (response.BookDetail, error)
	FindByUser(string) ([]response.Book, error)
	Create(models.Book) error
	GetCover(int) (string, error)
	GetChapterAudioByBookID(int) ([]string, error)
	Update(models.Book) error
	Delete(int) error
}
