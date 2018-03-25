package http

import (
	"net/http"
	"reflect"

	"github.com/labstack/echo"
)

type UrlParam map[string]string
type QueryParams map[string]string
type GetBody func(interface{}) error

type Handler func(UrlParam, QueryParams, GetBody) interface{}

func WrapperEcho(fn Handler) echo.HandlerFunc {
	return func(c echo.Context) error {
		url := make(map[string]string)
		urlParams := c.ParamNames()
		for _, urlParam := range urlParams {
			url[urlParam] = c.Param(urlParam)
		}

		query := make(map[string]string)
		queryParams := reflect.ValueOf(c.QueryParams()).MapKeys()
		for _, queryParam := range queryParams {
			v := queryParam.String()
			query[v] = c.QueryParam(v)
		}

		body := func(i interface{}) error {
			return c.Bind(i)
		}

		return c.JSON(http.StatusOK, fn(url, query, body))
	}
}
