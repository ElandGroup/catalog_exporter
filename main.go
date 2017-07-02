package main

import (
	"flag"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
	configutil "github.com/pangpanglabs/goutils/config"

	"github.com/elandgroup/catalog_exporter/config"
	"github.com/elandgroup/catalog_exporter/controllers"
	"github.com/elandgroup/catalog_exporter/filters"
)

func main() {
	appEnv := flag.String("app-env", os.Getenv("APP_ENV"), "app env")
	flag.Parse()

	var c config.Config
	if err := configutil.Read(*appEnv, &c); err != nil {
		panic(err)
	}

	e := echo.New()

	controllers.HomeController{}.Init(e.Group("/"))
	controllers.DiscountController{}.Init(e.Group("/discounts"))
	controllers.DiscountApiController{}.Init(e.Group("/api/discounts"))

	e.Static("/static", "static")
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())
	e.Use(filters.DbContext(c.Database))
	e.Use(filters.Logger(*appEnv))
	e.Use(filters.Tracer(c.Trace))

	e.Renderer = filters.NewTemplate()
	e.Validator = &filters.Validator{}
	e.Debug = c.Debug

	if err := e.Start(":" + c.Httpport); err != nil {
		log.Println(err)
	}

}
