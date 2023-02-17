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
	e.GET("isopen", isopen)                       //working
	e.POST("setopen", setopen)                    //working
	e.POST("setclosed", setclosed)                //working
	e.GET("customercount", customercount)         //working
	e.POST("customerenters", customerenters)      //working
	e.POST("customerleaves", customerleaves)      //working
	e.POST("customerreset", customerreset)        //working
	e.GET("listproducts", listproducts)           //working
	e.GET("getproducts", getproducts)             //working
	e.GET("getproduct", getproduct)               //working
	e.GET("getproductids", getproductids)         //working
	e.POST("addproduct", addproduct)              //working
	e.DELETE("removeproduct", removeproduct)      //working
	e.GET("getnotificitations", getnotifications) //expected to work
	e.POST("addnotification", addnotification)    //working

	e.Logger.Fatal(e.Start(":1312"))
}
