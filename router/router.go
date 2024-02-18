package router

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Run(port string) {

	e := echo.New()

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if he, ok := err.(*echo.HTTPError); ok {
			e.Logger.Error(he)
			c.String(he.Code, fmt.Sprintf("%v: %v", he.Code, he.Message))
		}
	}

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	e.Start(port)

}
