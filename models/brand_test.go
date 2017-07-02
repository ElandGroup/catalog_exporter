package models_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/pangpanglabs/goutils/test"

	"catalog_exporter/models"
)

func TestBrandCreate(t *testing.T) {
	d1 := models.Brand{
		Name: "name1",
		Code: "code1",
	}
	affected, err := d1.Create(ctx)
	test.Ok(t, err)
	test.Equals(t, affected, int64(1))
	test.Equals(t, d1.Id, int64(1))
	test.Equals(t, d1.CreatedAt.Format("2006-01-02"), time.Now().Format("2006-01-02"))
	test.Equals(t, d1.UpdatedAt.Format("2006-01-02"), time.Now().Format("2006-01-02"))

	d2 := models.Brand{
		Name: "name2",
		Code: "code2",
	}
	affected, err = d2.Create(ctx)
	test.Ok(t, err)
	test.Equals(t, affected, int64(1))
	test.Equals(t, d2.Id, int64(2))
	test.Equals(t, d1.CreatedAt.Format("2006-01-02"), time.Now().Format("2006-01-02"))
	test.Equals(t, d1.UpdatedAt.Format("2006-01-02"), time.Now().Format("2006-01-02"))
}

func TestBrandGetAndUpdate(t *testing.T) {
	d, err := models.Brand{}.GetById(ctx, 1)
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

func TestBrandGetAll(t *testing.T) {
	totalCount, items, err := models.Brand{}.GetAll(ctx, []string{"name"}, []string{"code"}, 0, 10)
	test.Ok(t, err)
	test.Equals(t, totalCount, int64(2))
	test.Equals(t, items[0].Id, int64(2))
	test.Equals(t, items[1].Id, int64(1))
}

func TestXXX(t *testing.T) {
	at, err := time.Parse("2006-01-02", "2017-12-31")
	test.Ok(t, err)
	test.Equals(t, at.Year(), 2017)
	test.Assert(t, at.Month() == 12, "Month should be equals to 12")
	fmt.Println(at)
}
