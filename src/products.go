package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dchest/uniuri"
	"github.com/labstack/echo/v4"
)

func getproductids(c echo.Context) error {
	var products = []string{}
	files, err := os.ReadDir("products/")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		prdct := fmt.Sprint(f.Name())
		products = append(products, prdct)
	}

	return c.String(http.StatusOK, strings.Join(products, ","))
}

func addproduct(c echo.Context) error {
	p := new(addProduct)
	if err := c.Bind(p); err != nil {
		return err
	}
	if p.TITLE == "" || p.PRIZE == "" || p.DESCRIPTION == "" || p.CALORIES == "" || p.IMAGE == "" {
		return c.String(http.StatusOK, "incomplete data. Missing something?")
	}
	storeProduct(p.TITLE, p.PRIZE, p.ALLERGENIC, p.DESCRIPTION, p.CALORIES, p.SALE, p.IMAGE)
	return c.String(http.StatusOK, "success")
}

func changeproduct(c echo.Context) error {
	p := new(changeProduct)
	if err := c.Bind(p); err != nil {
		return err
	}
	if p.ID == "" || p.KEY == "" || p.VALUE == "" {
		return c.String(http.StatusOK, "incomplete data. Missing something?")
	}
	dir := "products/" + p.ID
	changeKeyUnsafe(dir, p.KEY, p.VALUE)
	return c.String(http.StatusOK, "success")
}

func getproduct(c echo.Context) error {
	p := new(getProduct)
	if err := c.Bind(p); err != nil {
		return err
	}
	if _, err := os.Stat("products/" + string(p.ID)); errors.Is(err, os.ErrNotExist) {
		return c.String(http.StatusOK, ("invalid ID"))
	}
	var product = []string{}
	prdct := fmt.Sprint(p.ID + ": " + readKeyUnsafe("name", "products/"+p.ID+"/") + ";")
	product = append(product, prdct)
	prdct = fmt.Sprintln(p.ID + ": " + readKeyUnsafe("title", "products/"+p.ID+"/") + ";")
	product = append(product, prdct)
	prdct = fmt.Sprintln(p.ID + ": " + readKeyUnsafe("prize", "products/"+p.ID+"/") + ";")
	product = append(product, prdct)
	if readKeyUnsafe("allergenic", "products/"+p.ID+"/") != "" {
		prdct = fmt.Sprintln(p.ID + ": " + readKeyUnsafe("allergenic", "products/"+p.ID+"/") + ";")
		product = append(product, prdct)
	} else {
		prdct = fmt.Sprintln(p.ID + ": " + "nv" + ";")
		product = append(product, prdct)
	}
	prdct = fmt.Sprintln(p.ID + ": " + readKeyUnsafe("description", "products/"+p.ID+"/") + ";")
	product = append(product, prdct)
	return c.String(http.StatusOK, strings.Join(product, ""))
}

func removeproduct(c echo.Context) error {
	p := new(getProduct)
	if err := c.Bind(p); err != nil {
		return err
	}
	if _, err := os.Stat("products/" + string(p.ID)); errors.Is(err, os.ErrNotExist) {
		return c.String(http.StatusOK, ("invalid ID"))
	}
	_ = os.RemoveAll("products/" + p.ID)
	return c.String(http.StatusOK, ("removed"))
}

func getproducts(c echo.Context) error {
	var products = []string{}
	files, err := os.ReadDir("products/")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		var product = []string{}
		prdct := fmt.Sprint(f.Name() + ": " + readKeyUnsafe("name", "products/"+f.Name()+"/") + ";")
		product = append(product, prdct)
		prdct = fmt.Sprintln(f.Name() + ": " + readKeyUnsafe("title", "products/"+f.Name()+"/") + ";")
		product = append(product, prdct)
		prdct = fmt.Sprintln(f.Name() + ": " + readKeyUnsafe("prize", "products/"+f.Name()+"/") + ";")
		product = append(product, prdct)
		prdct = fmt.Sprintln(f.Name() + ": " + readKeyUnsafe("calories", "products/"+f.Name()+"/") + ";")
		product = append(product, prdct)
		prdct = fmt.Sprintln(f.Name() + ": " + readKeyUnsafe("image", "products/"+f.Name()+"/") + ";")
		product = append(product, prdct)

		if readKeyUnsafe("allergenic", "products/"+f.Name()+"/") != "" {
			prdct = fmt.Sprintln(f.Name() + ": " + readKeyUnsafe("allergenic", "products/"+f.Name()+"/") + ";")
			product = append(product, prdct)
		} else {
			prdct = fmt.Sprintln(f.Name() + ": " + "nv" + ";")
			product = append(product, prdct)
		}

		if readKeyUnsafe("sale", "products/"+f.Name()+"/") != "" {
			prdct = fmt.Sprintln(f.Name() + ": " + readKeyUnsafe("sale", "products/"+f.Name()+"/") + ";")
			product = append(product, prdct)
		} else {
			prdct = fmt.Sprintln(f.Name() + ": " + "nv" + ";")
			product = append(product, prdct)
		}

		prdct = fmt.Sprintln(f.Name() + ": " + readKeyUnsafe("description", "products/"+f.Name()+"/") + ";")
		product = append(product, prdct)
		joined := strings.Join(product, "")
		products = append(products, joined)
	}
	return c.String(http.StatusOK, strings.Join(products, "|\n"))
}

func listproducts(c echo.Context) error {
	var products = []string{}
	files, err := os.ReadDir("products/")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		prdct := fmt.Sprint(f.Name() + ": " + readKeyUnsafe("name", "products/"+f.Name()+"/"))
		products = append(products, prdct)
	}

	return c.String(http.StatusOK, strings.Join(products, ""))
}

func storeProduct(title string, prize string, allergenic string, description string, calories string, sale string, image string) {
	id := genIDproduct()

	createDir("products/" + id)

	addKeyUnsafe("title", title, "products/"+id)
	addKeyUnsafe("prize", prize, "products/"+id)
	addKeyUnsafe("name", genName(title), "products/"+id)
	addKeyUnsafe("image", image, "products/"+id)
	addKeyUnsafe("calories", calories, "products/"+id)
	if allergenic != "" {
		addKeyUnsafe("allergenic", allergenic, "products/"+id)
	}

	if sale != "" {
		addKeyUnsafe("sale", sale, "products/"+id)
	}

	addKeyUnsafe("description", description, "products/"+id)

}

func genIDproduct() string {
	res := uniuri.NewLen(5)
	if _, err := os.Stat("products/" + res); !os.IsNotExist(err) {
		res = genIDproduct()
	}
	return res
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
