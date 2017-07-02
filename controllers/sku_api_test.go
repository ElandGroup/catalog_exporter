package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"

	"github.com/elandgroup/catalog_exporter/models"

	"github.com/pangpanglabs/goutils/test"
)

func Test_SkuApiController_Create(t *testing.T) {
	req := httptest.NewRequest(echo.POST, "/api/skus", strings.NewReader(`{"name":"sku name", "desc":"sku desc", "startAt":"2017-01-01","endAt":"2017-12-31","actionType":"Percentage","skuAmount":10,"enable":true}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	test.Ok(t, handleWithFilter(SkuApiController{}.Create, echoApp.NewContext(req, rec)))
	test.Equals(t, http.StatusOK, rec.Code)

	var v struct {
		Result  models.Sku `json:"result"`
		Success bool       `json:"success"`
	}
	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	test.Equals(t, v.Result.Name, "sku name")
	test.Equals(t, v.Result.UpdatedAt.Format("2006-01-02"), "2017-01-01")
}

func Test_SkuApiController_Create2(t *testing.T) {
	req := httptest.NewRequest(echo.POST, "/api/skus", strings.NewReader(`{"name":"sku name#2", "desc":"sku desc#2", "startAt":"2017-02-01","endAt":"2017-11-30","actionType":"Percentage","skuAmount":20,"enable":true}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	test.Ok(t, handleWithFilter(SkuApiController{}.Create, echoApp.NewContext(req, rec)))
	test.Equals(t, http.StatusOK, rec.Code)

	var v struct {
		Result  models.Sku `json:"result"`
		Success bool       `json:"success"`
	}
	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	test.Equals(t, v.Result.Name, "sku name#2")
	test.Equals(t, v.Result.UpdatedAt.Format("2006-01-02"), "2017-02-01")
}

func Test_SkuApiController_Update(t *testing.T) {
	req := httptest.NewRequest(echo.PUT, "/api/skus/1", strings.NewReader(`{"name":"sku name2", "desc":"sku desc2", "startAt":"2017-01-02","endAt":"2017-12-30","actionType":"Percentage","skuAmount":10,"enable":true}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echoApp.NewContext(req, rec)
	c.SetPath("/api/skus/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	test.Ok(t, handleWithFilter(SkuApiController{}.Update, c))
	test.Equals(t, http.StatusOK, rec.Code)

	var v struct {
		Result  models.Sku `json:"result"`
		Success bool       `json:"success"`
	}
	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	test.Equals(t, v.Result.Name, "sku name2")
	test.Equals(t, v.Result.UpdatedAt.Format("2006-01-02"), "2017-01-02")
}

func Test_SkuApiController_GetOne(t *testing.T) {
	req := httptest.NewRequest(echo.GET, "/api/skus/1", nil)
	rec := httptest.NewRecorder()
	c := echoApp.NewContext(req, rec)
	c.SetPath("/api/skus/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	test.Ok(t, handleWithFilter(SkuApiController{}.GetOne, c))
	test.Equals(t, http.StatusOK, rec.Code)

	var v struct {
		Result  models.Sku `json:"result"`
		Success bool       `json:"success"`
	}
	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	test.Equals(t, v.Result.Name, "sku name2")
	test.Equals(t, v.Result.UpdatedAt.Format("2006-01-02"), "2017-01-02")
}

func Test_SkuApiController_GetAll_SortByAsc(t *testing.T) {
	req := httptest.NewRequest(echo.GET, "/api/skus?sortby=sku_amount&order=asc", nil)
	rec := httptest.NewRecorder()
	c := echoApp.NewContext(req, rec)
	test.Ok(t, handleWithFilter(SkuApiController{}.GetAll, c))
	test.Equals(t, http.StatusOK, rec.Code)

	var v struct {
		Result struct {
			TotalCount int
			Items      []models.Sku
		} `json:"result"`
		Success bool `json:"success"`
	}
	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	test.Equals(t, v.Result.TotalCount, 2)
	test.Equals(t, v.Result.Items[0].SkuId, int64(10))
}

func Test_SkuApiController_GetAll_SortByDesc(t *testing.T) {
	req := httptest.NewRequest(echo.GET, "/api/skus?sortby=sku_amount&order=desc", nil)
	rec := httptest.NewRecorder()
	c := echoApp.NewContext(req, rec)
	test.Ok(t, handleWithFilter(SkuApiController{}.GetAll, c))
	// fmt.Println(rec.Body.String())
	test.Equals(t, http.StatusOK, rec.Code)

	var v struct {
		Result struct {
			TotalCount int
			Items      []models.Sku
		} `json:"result"`
		Success bool `json:"success"`
	}
	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	test.Equals(t, v.Result.TotalCount, 2)
	test.Equals(t, v.Result.Items[0].SkuId, int64(20))
}
