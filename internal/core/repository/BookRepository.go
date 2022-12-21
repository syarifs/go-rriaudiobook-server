package repository

import (
	"go-rriaudiobook-server/internal/core/entity/models"
	"go-rriaudiobook-server/internal/core/entity/response"
)

type BookRepository interface {
	FindAll() ([]response.Book, error)
	FindByID(int) (response.BookDetail, error)
	CreateBook(models.Book) error
	InsertChapter(models.Chapter) error
	GetBookCover(int) (string, error)
	GetChapterAudio(int, int) (string, error)
	GetChapterAudioByBookID(int) ([]string, error)
	UpdateBook(models.Book) error
	DeleteBook(int) error
	UpdateChapter(models.Chapter) error
	DeleteChapter(int, int) error
}
