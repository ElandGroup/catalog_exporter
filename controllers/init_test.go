package controllers

import (
	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"

	"github.com/elandgroup/catalog_exporter/config"
	"github.com/elandgroup/catalog_exporter/filters"
)

var (
	echoApp          *echo.Echo
	handleWithFilter func(handlerFunc echo.HandlerFunc, c echo.Context) error
)

func init() {
	echoApp = echo.New()
	echoApp.Validator = &filters.Validator{}

	logger := filters.Logger("test")
	db := filters.DbContext(config.Database{Driver: "sqlite3", Connection: ":memory:"})

	handleWithFilter = func(handlerFunc echo.HandlerFunc, c echo.Context) error {
		return logger(db(handlerFunc))(c)
	}
}
