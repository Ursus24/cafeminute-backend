package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

var inshop = 0

func customerenters(c echo.Context) error {
	return c.String(http.StatusOK, "customer entered")
}
func customerleaves(c echo.Context) error {
	return c.String(http.StatusOK, "customer leaved")
}
func customercount(c echo.Context) error {
	return c.String(http.StatusOK, fmt.Sprint(inshop))
}
