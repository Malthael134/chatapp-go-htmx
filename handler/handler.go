package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/malthael134/chatapp-go-htmx/handler/api"
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

	api.Setup(e)

	e.Start(port)

}
