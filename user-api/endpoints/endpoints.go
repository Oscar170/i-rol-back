package endpoints

import (
	"net/http"
	"reflect"

	"github.com/Oscar170/i-rol-back/user-api/models"
	"github.com/labstack/echo"
)

var Users models.User

func SignIn(urlParams map[string]string, queryParams map[string]string, body map[string]interface{}) interface{} {

	return models.Status{
		Active: true,
	}
}

func SignOut(urlParams map[string]string, queryParams map[string]string, body map[string]interface{}) interface{} {
	return models.Status{
		Active: true,
	}
}

func SignUp(urlParams map[string]string, queryParams map[string]string, body map[string]interface{}) interface{} {
	return models.Status{
		Active: true,
	}
}

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
