package models

import (
	"context"
	"time"

	"github.com/go-xorm/xorm"

	"github.comsq/elandgroup/catalog_exporter/factory"
)

type Brand struct {
	Id   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`

	CreatedAt time.Time `json:"createdAt" xorm:"created"`
	UpdatedAt time.Time `json:"updatedAt" xorm:"updated"`
}

func (d *Brand) Create(ctx context.Context) (int64, error) {
	return factory.DB(ctx).Insert(d)
}
func (Brand) GetById(ctx context.Context, id int64) (*Brand, error) {
	var v Brand
	if has, err := factory.DB(ctx).ID(id).Get(&v); err != nil {
		return nil, err
	} else if !has {
		return nil, nil
	}
	return &v, nil
}
func (Brand) GetAll(ctx context.Context, sortby, order []string, offset, limit int) (totalCount int64, items []Brand, err error) {
	queryBuilder := func() *xorm.Session {
		q := factory.DB(ctx)
		if err := setSortOrder(q, sortby, order); err != nil {
			factory.Logger(ctx).Error(err)
		}
		return q
	}

	errc := make(chan error)
	go func() {
		v, err := queryBuilder().Count(&Brand{})
		if err != nil {
			errc <- err
			return
		}
		totalCount = v
		errc <- nil

	}()

	go func() {
		if err := queryBuilder().Limit(limit, offset).Find(&items); err != nil {
			errc <- err
			return
		}
		errc <- nil
	}()

	if err := <-errc; err != nil {
		return 0, nil, err
	}
	if err := <-errc; err != nil {
		return 0, nil, err
	}
	return
}
func (d *Brand) Update(ctx context.Context) (err error) {
	_, err = factory.DB(ctx).ID(d.Id).Update(d)
	return
}

func (Brand) Delete(ctx context.Context, id int64) (err error) {
	_, err = factory.DB(ctx).ID(id).Delete(&Brand{})
	return
}
