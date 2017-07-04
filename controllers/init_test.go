package controllers

import (
	"echosample/config"
	"flag"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"

	"github.comsq/elandgroup/catalog_exporter/filters"
)

var (
	echoApp          *echo.Echo
	handleWithFilter func(handlerFunc echo.HandlerFunc, c echo.Context) error
)

func init() {
	flag.Parse()
	echoApp = echo.New()
	echoApp.Validator = &filters.Validator{}

	logger := filters.Logger("test")
	//db := filters.DbContext(config.Database{Driver: "sqlite3", Connection: ":memory:"})
	db := filters.DbContext(config.Database{Driver: config.Config.Database.Driver, Connection: config.Config.Database.Connection})

	handleWithFilter = func(handlerFunc echo.HandlerFunc, c echo.Context) error {
		return logger(db(handlerFunc))(c)
	}
}
