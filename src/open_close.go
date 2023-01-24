package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var open = "false"

func isopen(c echo.Context) error {
	return c.String(http.StatusOK, open)
}
func setopen(c echo.Context) error {
	open = "true"
	return c.String(http.StatusOK, "set to open")
}
func setclosed(c echo.Context) error {
	open = "false"
	return c.String(http.StatusOK, "set to closed")
}
