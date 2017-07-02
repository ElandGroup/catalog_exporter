package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"

	"github.com/elandgroup/catalog_exporter/factory"
	"github.com/elandgroup/catalog_exporter/models"
)

type ContentController struct {
}

func (c ContentController) Init(g *echo.Group) {
	g.GET("", c.GetAll)
	g.GET("/new", c.New)
	g.POST("", c.Create)
	g.GET("/:id", c.GetOne)
	g.GET("/:id/edit", c.Edit)
	g.POST("/:id", c.Update)
}

func (ContentController) GetAll(c echo.Context) error {
	var v SearchInput
	if err := c.Bind(&v); err != nil {
		setFlashMessage(c, map[string]string{"warning": err.Error()})
	}
	if v.MaxResultCount == 0 {
		v.MaxResultCount = DefaultMaxResultCount
	}

	factory.Logger(c.Request().Context()).WithFields(logrus.Fields{
		"sortby":         v.Sortby,
		"order":          v.Order,
		"maxResultCount": v.MaxResultCount,
		"skipCount":      v.SkipCount,
	}).Info("SearchInput")

	totalCount, items, err := models.Content{}.GetAll(c.Request().Context(), v.Sortby, v.Order, v.SkipCount, v.MaxResultCount)
	if err != nil {
		return err
	}
	return c.Render(http.StatusOK, "content/index", map[string]interface{}{
		"TotalCount":     totalCount,
		"Contents":       items,
		"MaxResultCount": v.MaxResultCount,
	})
}
func (ContentController) New(c echo.Context) error {
	return c.Render(http.StatusOK, "content/new", map[string]interface{}{
		FlashName: getFlashMessage(c),
		"Form":    &models.Content{},
	})
}
func (ContentController) Create(c echo.Context) error {
	var v ContentInput
	if err := c.Bind(&v); err != nil {
		setFlashMessage(c, map[string]string{"error": err.Error()})
		return c.Redirect(http.StatusFound, "/content/new")
	}
	if err := c.Validate(&v); err != nil {
		setFlashMessage(c, map[string]string{"error": err.Error()})
		return c.Redirect(http.StatusFound, "/contents/new")
	}
	content, err := v.ToModel()
	if err != nil {
		setFlashMessage(c, map[string]string{"error": err.Error()})
		return c.Redirect(http.StatusFound, "/contents/new")
	}
	if _, err := content.Create(c.Request().Context()); err != nil {
		return err
	}
	return c.Redirect(http.StatusFound, fmt.Sprintf("/contents/%d", content.Id))
}
func (ContentController) GetOne(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.Render(http.StatusNotFound, "404", nil)
	}
	v, err := models.Content{}.GetById(c.Request().Context(), id)
	if err != nil {
		return err
	}
	if v == nil {
		return c.Render(http.StatusNotFound, "404", nil)
	}
	return c.Render(http.StatusOK, "content/show", map[string]interface{}{"Content": v})
}

func (ContentController) Edit(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.Render(http.StatusNotFound, "404", nil)
	}
	v, err := models.Content{}.GetById(c.Request().Context(), id)
	if err != nil {
		return err
	}
	if v == nil {
		return c.Render(http.StatusNotFound, "404", nil)
	}
	return c.Render(http.StatusOK, "content/edit", map[string]interface{}{
		FlashName: getFlashMessage(c),
		"Form":    v,
	})
}
func (ContentController) Update(c echo.Context) error {
	var v ContentInput
	if err := c.Bind(&v); err != nil {
		setFlashMessage(c, map[string]string{"error": err.Error()})
		return c.Redirect(http.StatusFound, "/content/new")
	}
	if err := c.Validate(&v); err != nil {
		setFlashMessage(c, map[string]string{"error": err.Error()})
		return c.Redirect(http.StatusFound, "/contents/new")
	}
	content, err := v.ToModel()
	if err != nil {
		setFlashMessage(c, map[string]string{"error": err.Error()})
		return c.Redirect(http.StatusFound, "/contents/new")
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.Render(http.StatusNotFound, "404", nil)
	}
	content.Id = id
	if err := content.Update(c.Request().Context()); err != nil {
		return err
	}
	return c.Redirect(http.StatusFound, fmt.Sprintf("/contents/%d", content.Id))
}
