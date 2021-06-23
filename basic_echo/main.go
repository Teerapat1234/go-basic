package main

import (
	"go_basic/basic_echo/internal"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		var response struct {
			ServerStatus string `json:"server_status"`
			Message      string `json:"message"`
		}
		response.ServerStatus = "Normal"
		response.Message = "Hello World"
		return c.JSON(http.StatusOK, response)
	})

	e.POST("/hello", internal.PrintValue)
	e.GET("/othercall/:id", internal.Call3rdSystem)

	e.Logger.Fatal(e.Start(":7778"))
}
