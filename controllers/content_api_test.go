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

func Test_ContentApiController_Create(t *testing.T) {
	req := httptest.NewRequest(echo.POST, "/api/contents", strings.NewReader(`{"name":"content name", "desc":"content desc", "startAt":"2017-01-01","endAt":"2017-12-31","actionType":"Percentage","contentAmount":10,"enable":true}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	test.Ok(t, handleWithFilter(ContentApiController{}.Create, echoApp.NewContext(req, rec)))
	test.Equals(t, http.StatusOK, rec.Code)

	var v struct {
		Result  models.Content `json:"result"`
		Success bool            `json:"success"`
	}
	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	test.Equals(t, v.Result.Name, "content name")
	test.Equals(t, v.Result.StartAt.Format("2006-01-02"), "2017-01-01")
}

func Test_ContentApiController_Create2(t *testing.T) {
	req := httptest.NewRequest(echo.POST, "/api/contents", strings.NewReader(`{"name":"content name#2", "desc":"content desc#2", "startAt":"2017-02-01","endAt":"2017-11-30","actionType":"Percentage","contentAmount":20,"enable":true}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	test.Ok(t, handleWithFilter(ContentApiController{}.Create, echoApp.NewContext(req, rec)))
	test.Equals(t, http.StatusOK, rec.Code)

	var v struct {
		Result  models.Content `json:"result"`
		Success bool            `json:"success"`
	}
	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	test.Equals(t, v.Result.Name, "content name#2")
	test.Equals(t, v.Result.StartAt.Format("2006-01-02"), "2017-02-01")
}

func Test_ContentApiController_Update(t *testing.T) {
	req := httptest.NewRequest(echo.PUT, "/api/contents/1", strings.NewReader(`{"name":"content name2", "desc":"content desc2", "startAt":"2017-01-02","endAt":"2017-12-30","actionType":"Percentage","contentAmount":10,"enable":true}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echoApp.NewContext(req, rec)
	c.SetPath("/api/contents/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	test.Ok(t, handleWithFilter(ContentApiController{}.Update, c))
	test.Equals(t, http.StatusOK, rec.Code)

	var v struct {
		Result  models.Content `json:"result"`
		Success bool            `json:"success"`
	}
	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	test.Equals(t, v.Result.Name, "content name2")
	test.Equals(t, v.Result.StartAt.Format("2006-01-02"), "2017-01-02")
}

func Test_ContentApiController_GetOne(t *testing.T) {
	req := httptest.NewRequest(echo.GET, "/api/contents/1", nil)
	rec := httptest.NewRecorder()
	c := echoApp.NewContext(req, rec)
	c.SetPath("/api/contents/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	test.Ok(t, handleWithFilter(ContentApiController{}.GetOne, c))
	test.Equals(t, http.StatusOK, rec.Code)

	var v struct {
		Result  models.Content `json:"result"`
		Success bool            `json:"success"`
	}
	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	test.Equals(t, v.Result.Name, "content name2")
	test.Equals(t, v.Result.StartAt.Format("2006-01-02"), "2017-01-02")
}

func Test_ContentApiController_GetAll_SortByAsc(t *testing.T) {
	req := httptest.NewRequest(echo.GET, "/api/contents?sortby=content_amount&order=asc", nil)
	rec := httptest.NewRecorder()
	c := echoApp.NewContext(req, rec)
	test.Ok(t, handleWithFilter(ContentApiController{}.GetAll, c))
	test.Equals(t, http.StatusOK, rec.Code)

	var v struct {
		Result struct {
			TotalCount int
			Items      []models.Content
		} `json:"result"`
		Success bool `json:"success"`
	}
	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	test.Equals(t, v.Result.TotalCount, 2)
	test.Equals(t, v.Result.Items[0].ContentAmount, float64(10))
}

func Test_ContentApiController_GetAll_SortByDesc(t *testing.T) {
	req := httptest.NewRequest(echo.GET, "/api/contents?sortby=content_amount&order=desc", nil)
	rec := httptest.NewRecorder()
	c := echoApp.NewContext(req, rec)
	test.Ok(t, handleWithFilter(ContentApiController{}.GetAll, c))
	// fmt.Println(rec.Body.String())
	test.Equals(t, http.StatusOK, rec.Code)

	var v struct {
		Result struct {
			TotalCount int
			Items      []models.Content
		} `json:"result"`
		Success bool `json:"success"`
	}
	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	test.Equals(t, v.Result.TotalCount, 2)
	test.Equals(t, v.Result.Items[0].ContentAmount, float64(20))
}
