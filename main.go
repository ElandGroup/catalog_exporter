package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	configutil "github.com/pangpanglabs/goutils/config"

	"github.comsq/elandgroup/catalog_exporter/config"
	"github.comsq/elandgroup/catalog_exporter/controllers"
	"github.comsq/elandgroup/catalog_exporter/filters"
)

func main() {
	appEnv := flag.String("app-env", os.Getenv("APP_ENV"), "app env")
	conn_env := flag.String("conn-env", os.Getenv("CONN_KEY"), "connection string")
	flag.Parse()

	var c config.Config
	if err := configutil.Read(*appEnv, &c); err != nil {
		panic(err)
	}
	c.Database.Connection = *conn_env
	fmt.Println(*conn_env)

	e := echo.New()

	controllers.HomeController{}.Init(e.Group("/"))
	controllers.ContentController{}.Init(e.Group("/contents"))
	controllers.ContentApiController{}.Init(e.Group("/api/contents"))

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
