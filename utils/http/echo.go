package http

import (
	"net/http"
	"reflect"

	"github.com/labstack/echo"
)

type Handler func(map[string]string, map[string]string, map[string]interface{}) interface{}

func WrapperEcho(fn Handler) echo.HandlerFunc {
	return func(c echo.Context) error {
		url := make(map[string]string)
		query := make(map[string]string)
		body := echo.Map{}

		urlParams := c.ParamNames()
		queryParams := reflect.ValueOf(c.QueryParams()).MapKeys()

		for _, urlParam := range urlParams {
			url[urlParam] = c.Param(urlParam)
		}

		for _, queryParam := range queryParams {
			v := queryParam.String()
			query[v] = c.QueryParam(v)
		}

		c.Bind(&body)

		return c.JSON(http.StatusOK, fn(url, query, body))
	}
}
