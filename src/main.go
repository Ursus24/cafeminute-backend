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
	e.GET("isopen", isopen)
	e.POST("setopen", setopen)
	e.POST("setclosed", setclosed)
	e.GET("customercount", customercount)
	e.POST("customerenters", customerenters)
	e.POST("customerleaves", customerleaves)
	e.POST("customerreset", customerreset)
	e.POST("addproduct", addproduct)
	e.GET("listproducts", listproducts)
	e.GET("getproducts", getproducts)
	e.GET("getproduct", getproduct)
	e.GET("addnotification", addnotification)
	e.GET("getproductids", getproductids)
	e.DELETE("removeproduct", removeproduct)

	e.Logger.Fatal(e.Start(":1323"))
}
