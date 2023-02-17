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

func addnotification(c echo.Context) error {
	n := new(addNotification)
	if err := c.Bind(n); err != nil {
		return err
	}
	if n.HEADING == "" || n.CONTENT == "" || n.DATE == "" || n.TIME == "" {
		return c.String(http.StatusOK, "incomplete data. Missing something?")
	}

	if checkDate(n.DATE) == false {
		return c.String(http.StatusOK, "date is not a date. check input")
	}

	if checkTime(n.TIME) == false {
		return c.String(http.StatusOK, "time is not a time. check input")
	}

	storeNotification(n.HEADING, n.CONTENT, n.DATE, n.TIME)
	return c.String(http.StatusOK, "success")
}

func changenotifications(c echo.Context) error {
	return c.String(http.StatusOK, "coming soon")
}

func getnotifications(c echo.Context) error {
	var products = []string{}
	files, err := os.ReadDir("notifications/")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		var product = []string{}
		prdct := fmt.Sprint(f.Name() + ": " + readKeyUnsafe("heading", "notifications/"+f.Name()+"/") + ";")
		product = append(product, prdct)
		prdct = fmt.Sprintln(f.Name() + ": " + readKeyUnsafe("content", "notifications/"+f.Name()+";"))
		product = append(product, prdct)
		prdct = fmt.Sprintln(f.Name() + ": " + readKeyUnsafe("date", "notifications/"+f.Name()+";"))
		product = append(product, prdct)
		prdct = fmt.Sprintln(f.Name() + ": " + readKeyUnsafe("time", "notifications/"+f.Name()+";"))
		product = append(product, prdct)
		joined := strings.Join(product, "")
		products = append(products, joined)
	}
	return c.String(http.StatusOK, strings.Join(products, "|"))
}

func removenotification(c echo.Context) error {
	p := new(getProduct)
	if err := c.Bind(p); err != nil {
		return err
	}
	if _, err := os.Stat("notifications/" + string(p.ID)); errors.Is(err, os.ErrNotExist) {
		return c.String(http.StatusOK, ("invalid ID"))
	}
	_ = os.RemoveAll("notifications/" + p.ID)
	return c.String(http.StatusOK, ("removed"))
}

func storeNotification(heading string, content string, date string, time string) {
	id := genIDnotifications()
	createDir("notifications/" + id)
	addKeyUnsafe("heading", heading, "notifications/"+id)
	addKeyUnsafe("content", content, "notifications/"+id)
	addKeyUnsafe("date", date, "notifications/"+id)
	addKeyUnsafe("time", time, "notifications/"+id)
}

func genIDnotifications() string {
	res := uniuri.NewLen(5)
	if _, err := os.Stat("notifications/" + res); !os.IsNotExist(err) {
		res = genIDnotifications()
	}
	return res
}
