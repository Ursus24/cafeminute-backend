package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"

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

	if !checkDate(n.DATE) {
		return c.String(http.StatusOK, "date is not a date. check input")
	}

	if !checkTime(n.TIME) {
		return c.String(http.StatusOK, "time is not a time. check input")
	}
	if NotificationExists(parseDate(n.DATE, n.TIME).String()[0:16]) == "nv" {
		storeNotification(n.HEADING, n.CONTENT, n.DATE, n.TIME, parseDate(n.DATE, n.TIME).String())
		return c.String(http.StatusOK, "success")
	}
	return c.String(http.StatusOK, "only one notification can be scheduled per minute")
}
func fetchnotification(c echo.Context) error {
	id := NotificationExists(time.Now().Format("2006-01-02 15:04"))
	var products = []string{}
	if id == "nv" {
		return c.String(http.StatusOK, "nv")
	} else {
		var product = []string{}
		prdct := fmt.Sprint(id + ": " + readKeyUnsafe("heading", "notifications/"+id+"/") + ";")
		product = append(product, prdct)
		prdct = fmt.Sprintln(id + ": " + readKeyUnsafe("content", "notifications/"+id+"/") + ";")
		product = append(product, prdct)
		prdct = fmt.Sprintln(id + ": " + readKeyUnsafe("date", "notifications/"+id+"/") + ";")
		product = append(product, prdct)
		prdct = fmt.Sprintln(id + ": " + readKeyUnsafe("time", "notifications/"+id+"/") + ";")
		product = append(product, prdct)
		prdct = fmt.Sprintln(id + ": " + parseDate(readKeyUnsafe("date", "notifications/"+id+"/"), readKeyUnsafe("time", "notifications/"+id+"/")).String() + ";")
		product = append(product, prdct)
		joined := strings.Join(product, "")
		products = append(products, joined)
		return c.String(http.StatusOK, strings.Join(products, "|"))
	}
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
		prdct = fmt.Sprintln(f.Name() + ": " + readKeyUnsafe("content", "notifications/"+f.Name()+"/") + ";")
		product = append(product, prdct)
		prdct = fmt.Sprintln(f.Name() + ": " + readKeyUnsafe("date", "notifications/"+f.Name()+"/") + ";")
		product = append(product, prdct)
		prdct = fmt.Sprintln(f.Name() + ": " + readKeyUnsafe("time", "notifications/"+f.Name()+"/") + ";")
		product = append(product, prdct)
		prdct = fmt.Sprintln(f.Name() + ": " + parseDate(readKeyUnsafe("date", "notifications/"+f.Name()+"/"), readKeyUnsafe("time", "notifications/"+f.Name()+"/")).String() + ";")
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

func storeNotification(heading string, content string, date string, time string, timefmt string) {
	id := genIDnotifications()
	createDir("notifications/" + id)
	addKeyUnsafe("heading", heading, "notifications/"+id)
	addKeyUnsafe("content", content, "notifications/"+id)
	addKeyUnsafe("date", date, "notifications/"+id)
	addKeyUnsafe("time", time, "notifications/"+id)
	addKeyUnsafe("timefmt", timefmt[0:16], "notifications/"+id)
}

func genIDnotifications() string {
	res := fmt.Sprint(genRandNum(10000, 99999))
	if _, err := os.Stat("notifications/" + res); !os.IsNotExist(err) {
		res = genIDnotifications()
	}
	return fmt.Sprint(res)
}

func genRandNum(min, max int64) int64 {
	// calculate the max we will be using
	bg := big.NewInt(max - min)

	// get big.Int between 0 and bg
	// in this case 0 to 20
	n, err := rand.Int(rand.Reader, bg)
	if err != nil {
		panic(err)
	}

	// add n to min to support the passed in range
	return n.Int64() + min
}

func NotificationExists(timeOfNot string) string {
	dates := []string{}
	IDs := []string{}
	files, err := os.ReadDir("notifications/")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		dates = append(dates, readKeyUnsafe("timefmt", "notifications/"+f.Name()))
		IDs = append(IDs, f.Name())
	}
	fmt.Println(timeOfNot)
	fmt.Println(dates)
	if contains(dates, timeOfNot) {
		return IDs[sort.StringSlice(dates).Search(timeOfNot)]
	}
	return "nv"
}
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
