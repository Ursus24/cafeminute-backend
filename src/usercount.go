package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var customer_level = 0

func getcustomers(c echo.Context) error {

	return c.String(http.StatusOK, fmt.Sprint(customer_level))
}
func setcustomers(c echo.Context) error {
	p := new(customers)
	if err := c.Bind(p); err != nil {
		return err
	}
	if p.PSWD == serverpswd {
		customer_level, _ = strconv.Atoi(p.CUSTOMERS)
		return c.String(http.StatusOK, "successs")
	}
	return c.String(http.StatusForbidden, "forbidden")
}
