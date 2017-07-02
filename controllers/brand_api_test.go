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

func Test_BrandApiController_Create(t *testing.T) {
	req := httptest.NewRequest(echo.POST, "/api/brands", strings.NewReader(`{"name":"brand name", "code":"code desc"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	test.Ok(t, handleWithFilter(BrandApiController{}.Create, echoApp.NewContext(req, rec)))
	test.Equals(t, http.StatusOK, rec.Code)

	var v struct {
		Result  models.Brand `json:"result"`
		Success bool         `json:"success"`
	}
	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	test.Equals(t, v.Result.Name, "brand name")
	test.Equals(t, v.Result.UpdatedAt.Format("2006-01-02"), "2017-01-01")
}

func Test_BrandApiController_Create2(t *testing.T) {
	req := httptest.NewRequest(echo.POST, "/api/brands", strings.NewReader(`{"name":"brand name#2", "code":"code desc#2"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	test.Ok(t, handleWithFilter(BrandApiController{}.Create, echoApp.NewContext(req, rec)))
	test.Equals(t, http.StatusOK, rec.Code)

	var v struct {
		Result  models.Brand `json:"result"`
		Success bool         `json:"success"`
	}
	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	test.Equals(t, v.Result.Name, "brand name#2")
	test.Equals(t, v.Result.UpdatedAt.Format("2006-01-02"), "2017-02-01")
}

func Test_BrandApiController_Update(t *testing.T) {
	req := httptest.NewRequest(echo.PUT, "/api/brands/1", strings.NewReader(`{"name":"brand name2", "desc":"brand desc2", "startAt":"2017-01-02","endAt":"2017-12-30","actionType":"Percentage","brandAmount":10,"enable":true}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echoApp.NewContext(req, rec)
	c.SetPath("/api/brands/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	test.Ok(t, handleWithFilter(BrandApiController{}.Update, c))
	test.Equals(t, http.StatusOK, rec.Code)

	var v struct {
		Result  models.Brand `json:"result"`
		Success bool         `json:"success"`
	}
	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	test.Equals(t, v.Result.Name, "brand name2")
	test.Equals(t, v.Result.UpdatedAt.Format("2006-01-02"), "2017-01-02")
}

func Test_BrandApiController_GetOne(t *testing.T) {
	req := httptest.NewRequest(echo.GET, "/api/brands/1", nil)
	rec := httptest.NewRecorder()
	c := echoApp.NewContext(req, rec)
	c.SetPath("/api/brands/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	test.Ok(t, handleWithFilter(BrandApiController{}.GetOne, c))
	test.Equals(t, http.StatusOK, rec.Code)

	var v struct {
		Result  models.Brand `json:"result"`
		Success bool         `json:"success"`
	}
	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	test.Equals(t, v.Result.Name, "brand name2")
	test.Equals(t, v.Result.UpdatedAt.Format("2006-01-02"), "2017-01-02")
}

func Test_BrandApiController_GetAll_SortByAsc(t *testing.T) {
	req := httptest.NewRequest(echo.GET, "/api/brands?sortby=brand_amount&order=asc", nil)
	rec := httptest.NewRecorder()
	c := echoApp.NewContext(req, rec)
	test.Ok(t, handleWithFilter(BrandApiController{}.GetAll, c))
	test.Equals(t, http.StatusOK, rec.Code)

	var v struct {
		Result struct {
			TotalCount int
			Items      []models.Brand
		} `json:"result"`
		Success bool `json:"success"`
	}
	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	test.Equals(t, v.Result.TotalCount, 2)
	test.Equals(t, v.Result.Items[0].Name, "brand name2")
}

func Test_BrandApiController_GetAll_SortByDesc(t *testing.T) {
	req := httptest.NewRequest(echo.GET, "/api/brands?sortby=brand_amount&order=desc", nil)
	rec := httptest.NewRecorder()
	c := echoApp.NewContext(req, rec)
	test.Ok(t, handleWithFilter(BrandApiController{}.GetAll, c))
	// fmt.Println(rec.Body.String())
	test.Equals(t, http.StatusOK, rec.Code)

	var v struct {
		Result struct {
			TotalCount int
			Items      []models.Brand
		} `json:"result"`
		Success bool `json:"success"`
	}
	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	test.Equals(t, v.Result.TotalCount, 2)
	test.Equals(t, v.Result.Items[0].Name, "brand name1")
}
