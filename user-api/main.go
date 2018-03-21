package main

import (
	"net/http"

	"github.com/Oscar170/i-rol-back/user-api/endpoints"
	"github.com/Oscar170/i-rol-back/user-api/models"

	"github.com/labstack/echo"
)

var (
	port   = ":8080"
	status = &models.Status{
		Active: true,
	}
)

func main() {
	e := echo.New()
	e.GET("/health-check", func(c echo.Context) error {
		return c.JSON(http.StatusOK, status)
	})

	e.POST("/sign-in", endpoints.WrapperEcho(endpoints.SignIn))
	e.GET("/sign-out", endpoints.WrapperEcho(endpoints.SignOut))
	e.GET("/sign-up", endpoints.WrapperEcho(endpoints.SignUp))

	e.Logger.Fatal(e.Start(port))
}
