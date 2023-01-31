package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	_ = db
	db.AutoMigrate(&Film{})
	db.Create(&Film{Name: "", IMDB: 12})
	var film Film
	db.First(&film, 1) // find product with integer primary key
	fmt.Println(film)
	if err != nil {
		panic("failed to connect database")
	}
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
	e.POST("customerReset", customerreset)
	e.POST("addProduct", addproduct)
	e.POST("listProducts", listproducts)
	e.POST("getProducts", getproducts)
	e.Logger.Fatal(e.Start(":1323"))
}
