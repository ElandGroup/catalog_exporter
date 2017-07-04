package models_test

import (
	"github.comsq/elandgroup/catalog_exporter/config"
	"context"
	"runtime"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"

	"github.comsq/elandgroup/catalog_exporter/factory"
	"github.comsq/elandgroup/catalog_exporter/models"
)

var ctx context.Context

func init() {
	runtime.GOMAXPROCS(1)
	xormEngine, err := xorm.NewEngine(config.Config.Database.Driver, config.Config.Database.Connection)
	if err != nil {
		panic(err)
	}
	xormEngine.ShowSQL(true)
	xormEngine.Sync(new(models.Content))
	ctx = context.WithValue(context.Background(), factory.ContextDBName, xormEngine.NewSession())
}
