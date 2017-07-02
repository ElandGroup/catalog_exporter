package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"

	"catalog_exporter/factory"
	"catalog_exporter/models"
)

type SkuController struct {
}

func (c SkuController) Init(g *echo.Group) {
	g.GET("", c.GetAll)
	g.GET("/new", c.New)
	g.POST("", c.Create)
	g.GET("/:id", c.GetOne)
	g.GET("/:id/edit", c.Edit)
	g.POST("/:id", c.Update)
}

func (SkuController) GetAll(c echo.Context) error {
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

	totalCount, items, err := models.Sku{}.GetAll(c.Request().Context(), v.Sortby, v.Order, v.SkipCount, v.MaxResultCount)
	if err != nil {
		return err
	}
	return c.Render(http.StatusOK, "sku/index", map[string]interface{}{
		"TotalCount":     totalCount,
		"Skus":           items,
		"MaxResultCount": v.MaxResultCount,
	})
}
func (SkuController) New(c echo.Context) error {
	return c.Render(http.StatusOK, "sku/new", map[string]interface{}{
		FlashName: getFlashMessage(c),
		"Form":    &models.Sku{},
	})
}
func (SkuController) Create(c echo.Context) error {
	var v SkuInput
	if err := c.Bind(&v); err != nil {
		setFlashMessage(c, map[string]string{"error": err.Error()})
		return c.Redirect(http.StatusFound, "/sku/new")
	}
	if err := c.Validate(&v); err != nil {
		setFlashMessage(c, map[string]string{"error": err.Error()})
		return c.Redirect(http.StatusFound, "/skus/new")
	}
	sku, err := v.ToModel()
	if err != nil {
		setFlashMessage(c, map[string]string{"error": err.Error()})
		return c.Redirect(http.StatusFound, "/skus/new")
	}
	if _, err := sku.Create(c.Request().Context()); err != nil {
		return err
	}
	return c.Redirect(http.StatusFound, fmt.Sprintf("/skus/%d", sku.Id))
}
func (SkuController) GetOne(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.Render(http.StatusNotFound, "404", nil)
	}
	v, err := models.Sku{}.GetById(c.Request().Context(), id)
	if err != nil {
		return err
	}
	if v == nil {
		return c.Render(http.StatusNotFound, "404", nil)
	}
	return c.Render(http.StatusOK, "sku/show", map[string]interface{}{"Sku": v})
}

func (SkuController) Edit(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.Render(http.StatusNotFound, "404", nil)
	}
	v, err := models.Sku{}.GetById(c.Request().Context(), id)
	if err != nil {
		return err
	}
	if v == nil {
		return c.Render(http.StatusNotFound, "404", nil)
	}
	return c.Render(http.StatusOK, "sku/edit", map[string]interface{}{
		FlashName: getFlashMessage(c),
		"Form":    v,
	})
}
func (SkuController) Update(c echo.Context) error {
	var v SkuInput
	if err := c.Bind(&v); err != nil {
		setFlashMessage(c, map[string]string{"error": err.Error()})
		return c.Redirect(http.StatusFound, "/sku/new")
	}
	if err := c.Validate(&v); err != nil {
		setFlashMessage(c, map[string]string{"error": err.Error()})
		return c.Redirect(http.StatusFound, "/skus/new")
	}
	sku, err := v.ToModel()
	if err != nil {
		setFlashMessage(c, map[string]string{"error": err.Error()})
		return c.Redirect(http.StatusFound, "/skus/new")
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.Render(http.StatusNotFound, "404", nil)
	}
	sku.Id = id
	if err := sku.Update(c.Request().Context()); err != nil {
		return err
	}
	return c.Redirect(http.StatusFound, fmt.Sprintf("/skus/%d", sku.Id))
}
