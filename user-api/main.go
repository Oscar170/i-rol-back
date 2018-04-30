package main

import (
	"net/http"

	"github.com/Oscar170/i-rol-back/user-api/endpoints"
	"github.com/Oscar170/i-rol-back/user-api/models"
	utilHttp "github.com/Oscar170/i-rol-back/utils/http"
	"github.com/labstack/echo"
)

var (
	port   = ":8080"
	health = &models.Health{
		Active: true,
	}
)

func main() {
	e := echo.New()

	e.GET("/health-check", func(c echo.Context) error {
		return c.JSON(http.StatusOK, health)
	})

	e.POST("/sign-up", utilHttp.WrapperEcho(endpoints.SignUp))
	e.POST("/sign-in", utilHttp.WrapperEcho(endpoints.SignIn))

	e.Logger.Fatal(e.Start(port))
}
