package models

import (
	"context"
	"time"

	"github.com/go-xorm/xorm"

	"github.com/elandgroup/catalog_exporter/factory"
)

type Content struct {
	Id        int64   `json:"id"`
	Code      string  `json:"code"`
	Name      string  `json:"name"`
	Desc      string  `json:"desc"`
	ListPrice float64 `json:"list_price"`

	Images    string    `json:"images"`
	Rank      string    `json:"rank"`
	BrandId   int64     `json:"brand_id"`
	Enable    bool      `json:"enable"`
	CreatedAt time.Time `json:"createdAt" xorm:"created"`
	UpdatedAt time.Time `json:"updatedAt" xorm:"updated"`
}

func (d *Content) Create(ctx context.Context) (int64, error) {
	return factory.DB(ctx).Insert(d)
}
func (Content) GetById(ctx context.Context, id int64) (*Content, error) {
	var v Content
	if has, err := factory.DB(ctx).ID(id).Get(&v); err != nil {
		return nil, err
	} else if !has {
		return nil, nil
	}
	return &v, nil
}
func (Content) GetAll(ctx context.Context, sortby, order []string, offset, limit int) (totalCount int64, items []Content, err error) {
	queryBuilder := func() *xorm.Session {
		q := factory.DB(ctx)
		if err := setSortOrder(q, sortby, order); err != nil {
			factory.Logger(ctx).Error(err)
		}
		return q
	}

	errc := make(chan error)
	go func() {
		v, err := queryBuilder().Count(&Content{})
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
func (d *Content) Update(ctx context.Context) (err error) {
	_, err = factory.DB(ctx).ID(d.Id).Update(d)
	return
}

func (Content) Delete(ctx context.Context, id int64) (err error) {
	_, err = factory.DB(ctx).ID(id).Delete(&Content{})
	return
}
