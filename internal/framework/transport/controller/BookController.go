package controller

import (
	"go-rriaudiobook-server/internal/core/entity/request"
	"go-rriaudiobook-server/internal/core/entity/response"
	"go-rriaudiobook-server/internal/core/service"
	"go-rriaudiobook-server/internal/utils/errors/check"
	"go-rriaudiobook-server/internal/utils/file"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookController struct {
	srv *service.BookService
}

func NewBookController(srv *service.BookService) *BookController {
	return &BookController{srv: srv}
}

// godoc
// @Summary GetAllBook
// @Description Fetch All Book Data
// @Tags Book
// @Security ApiKey
// @Accept json
// @Produce json
// @Success 200 {object} response.MessageData{data=[]response.Book}
// @Failure 417 {object} response.Error{}
// @Failure 500 {object} response.Error{}
// @Router /books [get]
func (bcon BookController) GetAllBook(c echo.Context) error {
	var res []response.Book
	var err error

	res, err = bcon.srv.FindAll()
	if r, ok := check.HTTP(res, err, "Fetch Book"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	return c.JSON(200, response.MessageData{
		Message: "Book Fetched",
		Data:    res,
	})
}

// godoc
// @Summary GetBookByID
// @Description Fetch Book Data By ID
// @Tags Book
// @Security ApiKey
// @Accept json
// @Produce json
// @Param id path int true "book id"
// @Success 200 {object} response.MessageData{data=response.BookDetail}
// @Failure 417 {object} response.Error{}
// @Failure 500 {object} response.Error{}
// @Router /books/{id} [get]
func (bcon BookController) GetBookByID(c echo.Context) error {
	var res response.BookDetail
	var err error
	var id, _ = strconv.Atoi(c.Param("id"))

	res, err = bcon.srv.FindByID(id)
	if r, ok := check.HTTP(res, err, "Fetch Book"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	return c.JSON(200, response.MessageData{
		Message: "Book Fetched",
		Data:    res,
	})
}

// godoc
// @Summary GetBookByID
// @Description Fetch Book Data By ID
// @Tags Book
// @Security ApiKey
// @Accept json
// @Produce json
// @Param id path int true "book id"
// @Success 200 {object} response.MessageData{data=response.BookDetail}
// @Failure 417 {object} response.Error{}
// @Failure 500 {object} response.Error{}
// @Router /books [post]
func (bcon BookController) Create(c echo.Context) error {
	var req request.BookRequest
	c.Bind(&req)
	cover, _ := c.FormFile("cover_image")
	req.CoverImage, _ = file.UploadFile(cover, "image/cover")

	if r, ok := check.HTTP(nil, req.Validate(), "Validate"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	err := bcon.srv.Create(req)
	if r, ok := check.HTTP(nil, err, "Create Book"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	return c.JSON(201, response.MessageOnly{
		Message: "Book Created",
	})
}
