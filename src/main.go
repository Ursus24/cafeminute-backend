package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, IP=${remote_ip}, status=${status}, latency:${latency},\n",
	}))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Root Route!")
	})
	e.GET("isOpen", isopen)
	e.POST("setOpen", setopen)
	e.POST("setClosed", setclosed)
	e.GET("customerCount", customercount)
	e.POST("customerEnters", customerenters)
	e.POST("customerLeaves", customerleaves)
	e.POST("customerReset", customerreset)
	e.POST("addProduct", addproduct)
	e.POST("listProducts", listproducts)
	e.POST("getProducts", getproducts)
	e.Logger.Fatal(e.Start(":1323"))
}
