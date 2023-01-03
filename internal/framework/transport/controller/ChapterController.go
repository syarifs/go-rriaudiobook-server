package controller

import (
	"go-rriaudiobook-server/internal/core/entity/request"
	"go-rriaudiobook-server/internal/core/entity/response"
	"go-rriaudiobook-server/internal/core/service"
	"go-rriaudiobook-server/internal/utils/errors/check"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ChapterController struct {
	srv *service.ChapterService
}

func NewChapterController(srv *service.ChapterService) *ChapterController {
	return &ChapterController{srv: srv}
}

// godoc
// @Summary Insert New Chapter
// @Description Insert New Chapter By Chapter ID
// @Tags Chapter
// @Security ApiKey
// @Accept json
// @Produce json
// @Param book_id path int true "book id"
// @Param body body request.ChapterRequest true "chapter data"
// @Success 200 {object} response.MessageOnly{}
// @Failure 417 {object} response.Error{}
// @Failure 500 {object} response.Error{}
// @Router /books/{book_id}/chapters [post]
func (bcon ChapterController) InsertChapter(c echo.Context) error {
	var req request.ChapterRequest
	var id, _ = strconv.Atoi(c.Param("book_id"))
	c.Bind(&req)
	req.BookID = uint(id)

	media_path, _ := c.FormFile("media_path")

	if media_path != nil {
		req.MediaPath = media_path
	}

	if r, ok := check.HTTP(nil, req.Validate(), "Validate"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	err := bcon.srv.Create(req)
	if r, ok := check.HTTP(nil, err, "Create Chapter"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	return c.JSON(201, response.MessageOnly{
		Message: "Chapter Created",
	})
}

// godoc
// @Summary Get All Chapter By Chapter ID
// @Description Fetch All Chapter Data By Chapter ID
// @Tags Chapter
// @Security ApiKey
// @Accept json
// @Produce json
// @Param book_id path int true "book id"
// @Param chapter_id path int true "book id"
// @Param body body request.ChapterRequest true "chapter data"
// @Success 200 {object} response.MessageData{data=response.ChapterDetail}
// @Failure 417 {object} response.Error{}
// @Failure 500 {object} response.Error{}
// @Router /books/{book_id}/chapters/{chapter_id}/update [put]
func (bcon ChapterController) UpdateChapter(c echo.Context) error {
	var req request.ChapterRequest
	var err error
	c.Bind(&req)
	var book_id, _ = strconv.Atoi(c.Param("book_id"))
	var chapter_id, _ = strconv.Atoi(c.Param("chapter_id"))
	req.BookID = uint(book_id)
	req.Code = chapter_id

	media_path, _ := c.FormFile("media_path")
	if media_path != nil {
		req.MediaPath = media_path
	}

	if r, ok := check.HTTP(nil, req.Validate(), "Validate"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	err = bcon.srv.Update(req)
	if r, ok := check.HTTP(req, err, "Update Chapter"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	return c.JSON(200, response.MessageOnly{
		Message: "Chapter Fetched",
	})
}

// godoc
// @Summary Delete Chapter By Chapter ID
// @Description Delete Chapter Data By ID
// @Tags Chapter
// @Security ApiKey
// @Accept json
// @Produce json
// @Param book_id path int true "book id"
// @Param chapter_id path int true "chapter id"
// @Success 200 {object} response.MessageOnly{}
// @Failure 417 {object} response.Error{}
// @Failure 500 {object} response.Error{}
// @Router /books/{book_id}/chapters/{chapter_id}/delete [delete]
func (bcon ChapterController) DeleteChapter(c echo.Context) error {
	var err error
	var book_id, _ = strconv.Atoi(c.Param("book_id"))
	var chapter_id, _ = strconv.Atoi(c.Param("chapter_id"))

	err = bcon.srv.Delete(book_id, chapter_id)
	if r, ok := check.HTTP(nil, err, "Delete Chapter"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	return c.JSON(200, response.MessageOnly{
		Message: "Chapter Deleted",
	})
}
