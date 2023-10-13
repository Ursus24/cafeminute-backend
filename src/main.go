package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var serverpswd = "123"

func main() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, IP=${remote_ip}, status=${status}, latency:${latency},\n",
	}))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Root Route!")
	})

	e.POST("setopen", setopen)           //working //secured //rusted
	e.POST("addschedule", add_schedule)  //working //secured //rusted
	e.GET("getschedule", get_schedule)   //working //rusted
	e.POST("setclosed", setclosed)       //working //secured //rusted
	e.GET("isopen", isopen)              //working //rusted
	e.POST("setcustomers", setcustomers) //working //secured //rusted
	e.GET("getcustomers", getcustomers)  //working //rusted
	e.POST("addproduct", addproduct)     //working //secured
	e.GET("listproducts", listproducts)  //working
	e.GET("getproducts", getproducts)    //working
	e.GET("getproduct", getproduct)      //working
	e.GET("getproductids", getproductids)
	e.PATCH("changeproduct", changeproduct)  //secured
	e.DELETE("removeproduct", removeproduct) //working //secured

	e.POST("addnotification", addnotification)         //working //secured
	e.GET("getnotifications", getnotifications)        //working
	e.GET("fetchnotification", fetchnotification)      //working
	e.PATCH("changenotification", changenotification)  //working //secured
	e.DELETE("removenotification", removenotification) //working //secured

	e.POST("addnews", addnews) //working //secured
	e.GET("getnewsids", getnewsids)
	e.GET("getnews", getnews)
	e.GET("getallnews", getallnews)
	e.GET("listnews", listnews)
	e.PATCH("changenews", changenews) //working //secured
	e.PATCH("addparticipant", addparticipant)
	e.PATCH("removeparticipant", removeparticipant)
	e.DELETE("removenews", removenews) //working //secured
	e.Logger.Fatal(e.Start(":13120"))
}

func genName(title string) string {
	res := fmt.Sprintln(strings.ReplaceAll(title, " ", ""))
	res = strings.ToLower(res)
	return res
}

func createDir(dir string) {
	path := dir
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
}
