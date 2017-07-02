package models_test

import (
	"testing"
	"time"

	"github.com/pangpanglabs/goutils/test"

	"github.com/elandgroup/catalog_exporter/models"
)

func TestSkuCreate(t *testing.T) {
	d1 := models.Sku{
		Name:    "name1",
		Options: "desc1",
	}
	affected, err := d1.Create(ctx)
	test.Ok(t, err)
	test.Equals(t, affected, int64(1))
	test.Equals(t, d1.Id, int64(1))
	test.Equals(t, d1.CreatedAt.Format("2006-01-02"), time.Now().Format("2006-01-02"))
	test.Equals(t, d1.UpdatedAt.Format("2006-01-02"), time.Now().Format("2006-01-02"))

	d2 := models.Sku{
		Name:    "name2",
		Options: "desc2",
	}
	affected, err = d2.Create(ctx)
	test.Ok(t, err)
	test.Equals(t, affected, int64(1))
	test.Equals(t, d2.Id, int64(2))
	test.Equals(t, d1.CreatedAt.Format("2006-01-02"), time.Now().Format("2006-01-02"))
	test.Equals(t, d1.UpdatedAt.Format("2006-01-02"), time.Now().Format("2006-01-02"))
}

func TestSkuGetAndUpdate(t *testing.T) {
	d, err := models.Sku{}.GetById(ctx, 1)
	test.Ok(t, err)
	test.Equals(t, d.Id, int64(1))
	test.Equals(t, d.Name, "name1")
	test.Equals(t, d.CreatedAt.Format("2006-01-02"), time.Now().Format("2006-01-02"))
	test.Equals(t, d.UpdatedAt.Format("2006-01-02"), time.Now().Format("2006-01-02"))

	d.Name = "name1-2"
	err = d.Update(ctx)
	test.Ok(t, err)
	test.Equals(t, d.Name, "name1-2")

}

func TestSkuGetAll(t *testing.T) {
	totalCount, items, err := models.Sku{}.GetAll(ctx, []string{"name"}, []string{"desc"}, 0, 10)
	test.Ok(t, err)
	test.Equals(t, totalCount, int64(2))
	test.Equals(t, items[0].Id, int64(2))
	test.Equals(t, items[1].Id, int64(1))
}
