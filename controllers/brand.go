package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"

	"github.comsq/elandgroup/catalog_exporter/factory"
	"github.comsq/elandgroup/catalog_exporter/models"
)

type BrandController struct {
}

func (c BrandController) Init(g *echo.Group) {
	g.GET("", c.GetAll)
	g.GET("/new", c.New)
	g.POST("", c.Create)
	g.GET("/:id", c.GetOne)
	g.GET("/:id/edit", c.Edit)
	g.POST("/:id", c.Update)
}

func (BrandController) GetAll(c echo.Context) error {
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

	totalCount, items, err := models.Brand{}.GetAll(c.Request().Context(), v.Sortby, v.Order, v.SkipCount, v.MaxResultCount)
	if err != nil {
		return err
	}
	return c.Render(http.StatusOK, "brand/index", map[string]interface{}{
		"TotalCount":     totalCount,
		"Brands":         items,
		"MaxResultCount": v.MaxResultCount,
	})
}
func (BrandController) New(c echo.Context) error {
	return c.Render(http.StatusOK, "brand/new", map[string]interface{}{
		FlashName: getFlashMessage(c),
		"Form":    &models.Brand{},
	})
}
func (BrandController) Create(c echo.Context) error {
	var v BrandInput
	if err := c.Bind(&v); err != nil {
		setFlashMessage(c, map[string]string{"error": err.Error()})
		return c.Redirect(http.StatusFound, "/brand/new")
	}
	if err := c.Validate(&v); err != nil {
		setFlashMessage(c, map[string]string{"error": err.Error()})
		return c.Redirect(http.StatusFound, "/brands/new")
	}
	brand, err := v.ToModel()
	if err != nil {
		setFlashMessage(c, map[string]string{"error": err.Error()})
		return c.Redirect(http.StatusFound, "/brands/new")
	}
	if _, err := brand.Create(c.Request().Context()); err != nil {
		return err
	}
	return c.Redirect(http.StatusFound, fmt.Sprintf("/brands/%d", brand.Id))
}
func (BrandController) GetOne(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.Render(http.StatusNotFound, "404", nil)
	}
	v, err := models.Brand{}.GetById(c.Request().Context(), id)
	if err != nil {
		return err
	}
	if v == nil {
		return c.Render(http.StatusNotFound, "404", nil)
	}
	return c.Render(http.StatusOK, "brand/show", map[string]interface{}{"Brand": v})
}

func (BrandController) Edit(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.Render(http.StatusNotFound, "404", nil)
	}
	v, err := models.Brand{}.GetById(c.Request().Context(), id)
	if err != nil {
		return err
	}
	if v == nil {
		return c.Render(http.StatusNotFound, "404", nil)
	}
	return c.Render(http.StatusOK, "brand/edit", map[string]interface{}{
		FlashName: getFlashMessage(c),
		"Form":    v,
	})
}
func (BrandController) Update(c echo.Context) error {
	var v BrandInput
	if err := c.Bind(&v); err != nil {
		setFlashMessage(c, map[string]string{"error": err.Error()})
		return c.Redirect(http.StatusFound, "/brand/new")
	}
	if err := c.Validate(&v); err != nil {
		setFlashMessage(c, map[string]string{"error": err.Error()})
		return c.Redirect(http.StatusFound, "/brands/new")
	}
	brand, err := v.ToModel()
	if err != nil {
		setFlashMessage(c, map[string]string{"error": err.Error()})
		return c.Redirect(http.StatusFound, "/brands/new")
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.Render(http.StatusNotFound, "404", nil)
	}
	brand.Id = id
	if err := brand.Update(c.Request().Context()); err != nil {
		return err
	}
	return c.Redirect(http.StatusFound, fmt.Sprintf("/brands/%d", brand.Id))
}
