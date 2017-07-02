package controllers

import "catalog_exporter/models"

const (
	DefaultMaxResultCount = 30
)

type SearchInput struct {
	Sortby         []string `query:"sortby"`
	Order          []string `query:"order"`
	SkipCount      int      `query:"skip_count"`
	MaxResultCount int      `query:"max_result_count"`
}
type ContentInput struct {
	Code      string  `json:"code" valid:"required"`
	Name      string  `json:"name" valid:"required"`
	Desc      string  `json:"desc"`
	ListPrice float64 `json:"list_price" valid:"required"`

	Images  string `json:"images"`
	Rank    string `json:"rank"`
	BrandId int64  `json:"brand_id" valid:"required"`
	Enable  bool   `json:"enable"`
}

func (d *ContentInput) ToModel() (*models.Content, error) {

	return &models.Content{
		Code:      d.Code,
		Name:      d.Name,
		Desc:      d.Desc,
		ListPrice: d.ListPrice,
		Images:    d.Images,

		Rank:    d.Rank,
		BrandId: d.BrandId,
		Enable:  d.Enable,
	}, nil
}

type SkuInput struct {
	SkuId int64  `json:"sku_idname" valid:"required"`
	Code  string `json:"code"`
	Name  string `json:"name" valid:"required"`

	Images  string `json:"images"`
	Options string `json:"options"`
	Enable  bool   `json:"enable"`
}

func (d *SkuInput) ToModel() (*models.Sku, error) {

	return &models.Sku{
		SkuId: d.SkuId,
		Code:  d.Code,
		Name:  d.Name,

		Images:  d.Images,
		Options: d.Options,
		Enable:  d.Enable,
	}, nil
}

type BrandInput struct {
	Code string `json:"code" valid:"required"`
	Name string `json:"name" valid:"required"`
}

func (d *BrandInput) ToModel() (*models.Brand, error) {

	return &models.Brand{
		Code: d.Code,
		Name: d.Name,
	}, nil
}
