package controller

import (
	"go-rriaudiobook-server/internal/core/entity/request"
	"go-rriaudiobook-server/internal/core/entity/response"
	"go-rriaudiobook-server/internal/core/service"
	"go-rriaudiobook-server/internal/utils/errors/check"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	srv *service.CategoryService
}

func NewCategoryController(srv *service.CategoryService) *CategoryController {
	return &CategoryController{srv: srv}
}

// godoc
// @Summary GetAllCategory
// @Description Fetch All Category Data
// @Tags Category
// @Security ApiKey
// @Accept json
// @Produce json
// @Success 200 {object} response.MessageData{data=[]response.Category}
// @Failure 417 {object} response.Error{}
// @Failure 500 {object} response.Error{}
// @Router /categories [get]
func (bcon CategoryController) FindAll(c echo.Context) error {
	var res []response.Category
	var err error

	res, err = bcon.srv.FindAll()
	if r, ok := check.HTTP(res, err, "Fetch Category"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	return c.JSON(200, response.MessageData{
		Message: "Category Fetched",
		Data:    res,
	})
}

// godoc
// @Summary Insert New Category
// @Description Fetch Category Data By ID
// @Tags Category
// @Security ApiKey
// @Accept json
// @Produce json
// @Param body body request.CategoryRequest true "book data"
// @Success 200 {object} response.MessageOnly{}
// @Failure 417 {object} response.Error{}
// @Failure 500 {object} response.Error{}
// @Router /categories [post]
func (bcon CategoryController) Create(c echo.Context) error {
	var req request.CategoryRequest
	c.Bind(&req)

	if r, ok := check.HTTP(nil, req.Validate(), "Validate"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	err := bcon.srv.Create(req)
	if r, ok := check.HTTP(nil, err, "Create Category"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	return c.JSON(201, response.MessageOnly{
		Message: "Category Created",
	})
}

// godoc
// @Summary UpdateCategoory
// @Description Update Category Data By ID
// @Tags Category
// @Security ApiKey
// @Accept json
// @Produce json
// @Param id path int true "book id"
// @Param body body request.CategoryRequest true "book data"
// @Success 200 {object} response.MessageOnly{}
// @Failure 417 {object} response.Error{}
// @Failure 500 {object} response.Error{}
// @Router /categories/{id}/update [put]
func (bcon CategoryController) Update(c echo.Context) error {
	var err error
	var req request.CategoryRequest
	var id, _ = strconv.Atoi(c.Param("id"))

	c.Bind(&req)
	req.ID = uint(id)

	err = bcon.srv.Update(req)
	if r, ok := check.HTTP(nil, err, "Update Category"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	return c.JSON(200, response.MessageOnly{
		Message: "Category Updated",
	})
}

// godoc
// @Summary DeleteCategoryByID
// @Description Delete Category Data By ID
// @Tags Category
// @Security ApiKey
// @Accept json
// @Produce json
// @Param id path int true "book id"
// @Success 200 {object} response.MessageOnly{}
// @Failure 417 {object} response.Error{}
// @Failure 500 {object} response.Error{}
// @Router /categories/{id} [delete]
func (bcon CategoryController) Delete(c echo.Context) error {
	var err error
	var id, _ = strconv.Atoi(c.Param("id"))

	err = bcon.srv.Delete(id)
	if r, ok := check.HTTP(nil, err, "Delete Category"); !ok {
		return c.JSON(r.Code, r.Result)
	}

	return c.JSON(200, response.MessageOnly{
		Message: "Category Deleted",
	})
}
