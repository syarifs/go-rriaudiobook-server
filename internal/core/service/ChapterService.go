package service

import (
	"go-rriaudiobook-server/internal/core/entity/models"
	"go-rriaudiobook-server/internal/core/entity/request"
	"go-rriaudiobook-server/internal/core/repository"
	"go-rriaudiobook-server/internal/utils"
	"go-rriaudiobook-server/internal/utils/errors"
	"go-rriaudiobook-server/internal/utils/file"
)

type ChapterService struct {
	repo repository.ChapterRepository
}

func NewChapterService(repo repository.ChapterRepository) *ChapterService {
	return &ChapterService{repo: repo}
}

func (bs ChapterService) Create(req request.ChapterRequest) (err error) {
	if req.BookID == 0 {
		return errors.New(400, "No ID Provided")
	}

	m, _ := utils.TypeConverter[models.Chapter](req)
	m.MediaPath, err = file.UploadFile(req.MediaPath, "chapter/", "audio/mpeg")
	if err != nil {
		return
	}

	err = bs.repo.Create(m)
	return
}

func (bs ChapterService) Update(req request.ChapterRequest) (err error) {
	if req.Code == 0 {
		return errors.New(400, "No ID Provided")
	}

	var media_file string
	media_file, err = bs.repo.GetAudio(req.Code, int(req.BookID))
	if err == nil {
		err = file.RemoveFile(media_file)
		if err != nil {
			return
		}
	}

	m, _ := utils.TypeConverter[models.Chapter](req)
	m.Code = uint(req.Code)
	m.MediaPath, err = file.UploadFile(req.MediaPath, "chapter/", "audio/mpeg")
	if err != nil {
		return
	}

	err = bs.repo.Update(m)
	return
}

func (bs ChapterService) Delete(book_id, chapter_id int) (err error) {
	if book_id == 0 {
		return errors.New(400, "No ID Provided")
	}

	var cover string
	cover, err = bs.repo.GetAudio(book_id, chapter_id)
	if err == nil {
		err = file.RemoveFile(cover)
		if err != nil {
			return
		}
	}

	err = bs.repo.Delete(chapter_id, book_id)
	return
}
