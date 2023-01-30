package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	_ "gorm.io/driver/sqlite"
	_ "gorm.io/gorm"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Root Route!")
	})
	e.GET("isOpen", isopen)
	e.POST("setOpen", setopen)
	e.POST("setClosed", setclosed)
	e.GET("customerCount", customercount)
	e.POST("customerEnters", customerenters)
	e.POST("customerLeaves", customerleaves)
	//test upload +1
	e.Logger.Fatal(e.Start(":1323"))
}
