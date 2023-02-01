package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"

	clover "github.com/ostafen/clover/v2"

	"github.com/labstack/echo/v4"
	//"gorm.io/driver/sqlite"
	//"gorm.io/gorm"
)

func addproduct(c echo.Context) error {
	i := new(addProduct)
	if err := c.Bind(i); err != nil {
		return err
	}

	if i.title == "" || i.prize == "" || i.descrition == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid data: Missing something?")
	}
	storeProduct(i.title, i.prize, i.allergenic, i.descrition)
	return c.String(http.StatusOK, "success")
}

func getproducts(c echo.Context) error {
	return c.String(http.StatusOK, "")
}

func listproducts(c echo.Context) error {
	db, _ := clover.Open("db")
	defer db.Close()
	docs, _ := db.FindAll(clover.NewQuery("products").Sort())
	var product = []string{}
	for _, doc := range docs {
		CERTAINPRODUCT := fmt.Sprintf("title: %s\n", doc.Get("title"), "id: %s\n", doc.Get("id"))
		product = append(product, CERTAINPRODUCT)
	}
	return c.String(http.StatusOK, strings.Join(product, "\n"))
}

func storeProduct(title string, prize string, allergenic string, description string) {
	db, _ := clover.Open("db")
	defer db.Close()
	doc := clover.NewDocument()
	doc.Set("id", generateID())
	doc.Set("title", title)
	doc.Set("name", getName(title))
	doc.Set("prize", prize)
	doc.Set("allergenic", allergenic)
	doc.Set("describtion", description)
	db.InsertOne("products", doc)
}

func getName(title string) string {
	result := strings.ReplaceAll(title, " ", "")
	result = strings.ToLower(result)
	return result
}

func generateID() string {
	db, _ := clover.Open("db")
	defer db.Close()
	id := RandStringBytes(4)
	docs, _ := db.FindAll(clover.NewQuery("products").Sort())
	for _, doc := range docs {
		idFromDB := fmt.Sprintf("%s\n", doc.Get(id))
		if idFromDB == id {
			generateID()
		}
	}
	return id
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
