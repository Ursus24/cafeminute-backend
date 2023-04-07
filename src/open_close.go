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
	p := new(setOpen)
	if err := c.Bind(p); err != nil {
		return err
	}
	if string(p.PSWD) == serverpswd {
		open = "true"
		return c.String(http.StatusOK, "sucess")
	}
	return c.String(http.StatusOK, "forbidden")
}

func setclosed(c echo.Context) error {
	p := new(setOpen)
	if err := c.Bind(p); err != nil {
		return err
	}
	if string(p.PSWD) == serverpswd {
		open = "false"
		return c.String(http.StatusOK, "success")
	}
	return c.String(http.StatusOK, "forbidden")
}
