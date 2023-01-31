package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	//"gorm.io/driver/sqlite"
	//"gorm.io/gorm"
)

func addproduct(c echo.Context) error {

	return c.String(http.StatusOK, "")
}

func getproducts(c echo.Context) error {
	return c.String(http.StatusOK, "")
}
func listproducts(c echo.Context) error {
	return c.String(http.StatusOK, "")
}
