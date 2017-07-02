package controllers

import "github.com/elandgroup/catalog_exporter/models"

const (
	DefaultMaxResultCount = 30
)

type SearchInput struct {
	Sortby         []string `query:"sortby"`
	Order          []string `query:"order"`
	SkipCount      int      `query:"skipCount"`
	MaxResultCount int      `query:"maxResultCount"`
}
type ContentInput struct {
	Code      string `json:"code" valid:"required"`
	Name      string `json:"name" valid:"required"`
	Desc      string `json:"desc"`
	ListPrice string `json:"list_price" valid:"required"`

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
