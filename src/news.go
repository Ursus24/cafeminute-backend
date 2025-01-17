package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dchest/uniuri"
	"github.com/labstack/echo/v4"
	stripmd "github.com/writeas/go-strip-markdown"
)

func addnews(c echo.Context) error {
	p := new(addNews)
	if err := c.Bind(p); err != nil {
		return err
	}
	if p.PSWD == serverpswd {
		if p.HEADING == "" || p.CONTENT == "" || p.IMAGE == "" || p.EVENT == "" && p.EVENT != "true" && p.EVENT != "false" {
			return c.String(http.StatusOK, "incomplete data. Missing something?")
		}
		storeNews(p.HEADING, p.CONTENT, p.IMAGE, p.EVENT)
		return c.String(http.StatusOK, "success")
	}
	return c.String(http.StatusForbidden, "forbidden")
}

func removenews(c echo.Context) error {
	p := new(removeNews)
	if err := c.Bind(p); err != nil {
		return err
	}
	if p.PSWD == serverpswd {
		if _, err := os.Stat("news/" + string(p.ID)); errors.Is(err, os.ErrNotExist) {
			return c.String(http.StatusOK, ("invalid ID"))
		}
		_ = os.RemoveAll("news/" + p.ID)
		return c.String(http.StatusOK, ("removed"))
	}
	return c.String(http.StatusOK, ("forbidden"))
}

func storeNews(heading string, content string, image string, event string) {
	id := genIDnews()
	createDir("news/" + id)
	addKeyUnsafe("heading", heading, "news/"+id)
	addKeyUnsafe("content", content, "news/"+id)
	addKeyUnsafe("contentRaw", stripmd.Strip(content), "news/"+id)
	addKeyUnsafe("image", image, "news/"+id)
	addKeyUnsafe("date", time.Now().Format("2006-01-02"), "news/"+id)
	addKeyUnsafe("event", event, "news/"+id)
	if event == "true" {
		addKeyUnsafe("participants", "0", "news/"+id)
	} else {
		addKeyUnsafe("participants", "nv", "news/"+id)
	}
}

func genIDnews() string {
	res := uniuri.NewLen(5)
	if _, err := os.Stat("news/" + res); !os.IsNotExist(err) {
		res = genIDnews()
	}
	return res
}

func getnewsids(c echo.Context) error {
	var news = []string{}
	files, err := os.ReadDir("news/")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		nws := fmt.Sprint(f.Name())
		news = append(news, nws)
	}

	return c.String(http.StatusOK, strings.Join(news, ","))
}

func getnews(c echo.Context) error {
	p := new(getNews)
	if err := c.Bind(p); err != nil {
		return err
	}
	if _, err := os.Stat("news/" + string(p.ID)); errors.Is(err, os.ErrNotExist) {
		return c.String(http.StatusOK, ("invalid ID"))
	}
	var news = []string{}
	nws := fmt.Sprint(p.ID + ": " + readKeyUnsafe("content", "news/"+p.ID+"/") + ";")
	news = append(news, nws)
	nws = fmt.Sprintln(p.ID + ": " + readKeyUnsafe("heading", "news/"+p.ID+"/") + ";")
	news = append(news, nws)
	nws = fmt.Sprintln(p.ID + ": " + readKeyUnsafe("contentRaw", "news/"+p.ID+"/") + ";")
	news = append(news, nws)
	nws = fmt.Sprintln(p.ID + ": " + readKeyUnsafe("image", "news/"+p.ID+"/") + ";")
	news = append(news, nws)
	nws = fmt.Sprintln(p.ID + ": " + readKeyUnsafe("date", "news/"+p.ID+"/") + ";")
	news = append(news, nws)
	return c.String(http.StatusOK, strings.Join(news, ""))
}
func getallnews(c echo.Context) error {
	var products = []string{}
	files, err := os.ReadDir("news/")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		var product = []string{}
		prdct := fmt.Sprint(f.Name() + ": " + readKeyUnsafe("content", "news/"+f.Name()+"/") + ";")
		product = append(product, prdct)
		prdct = fmt.Sprintln(f.Name() + ": " + readKeyUnsafe("contentRaw", "news/"+f.Name()+"/") + ";")
		product = append(product, prdct)
		prdct = fmt.Sprintln(f.Name() + ": " + readKeyUnsafe("date", "news/"+f.Name()+"/") + ";")
		product = append(product, prdct)
		prdct = fmt.Sprintln(f.Name() + ": " + readKeyUnsafe("heading", "news/"+f.Name()+"/") + ";")
		product = append(product, prdct)
		prdct = fmt.Sprintln(f.Name() + ": " + readKeyUnsafe("image", "news/"+f.Name()+"/") + ";")
		product = append(product, prdct)
		prdct = fmt.Sprintln(f.Name() + ": " + readKeyUnsafe("event", "news/"+f.Name()+"/") + ";")
		product = append(product, prdct)
		prdct = fmt.Sprintln(f.Name() + ": " + readKeyUnsafe("participants", "news/"+f.Name()+"/") + ";")
		product = append(product, prdct)
		joined := strings.Join(product, "")
		products = append(products, joined)
	}
	return c.String(http.StatusOK, strings.Join(products, "|\n"))
}

func listnews(c echo.Context) error {
	var news = []string{}
	files, err := os.ReadDir("news/")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		nws := fmt.Sprint(f.Name() + ": " + genName(readKeyUnsafe("heading", "news/"+f.Name()+"/")))
		news = append(news, nws)
	}
	return c.String(http.StatusOK, strings.Join(news, ""))
}

func changenews(c echo.Context) error {
	p := new(changeNews)
	if err := c.Bind(p); err != nil {
		return err
	}
	if p.PSWD == serverpswd {
		if p.ID == "" || p.KEY == "" || p.VALUE == "" {
			return c.String(http.StatusOK, "incomplete data. Missing something?")
		}
		dir := "news/" + p.ID
		changeKeyUnsafe(p.KEY, p.VALUE, dir)
		return c.String(http.StatusOK, "success")
	}

	return c.String(http.StatusForbidden, "forbidden")
}

func addparticipant(c echo.Context) error {

	p := new(removeNews)
	if err := c.Bind(p); err != nil {
		return err
	}

	if p.PSWD == serverpswd {

		old := readKeyUnsafe("participants", "news/"+p.ID+"/")
		fmt.Println(old)
		if old != "nv" {
			new, _ := strconv.Atoi(old)
			new = new + 1
			newstring := fmt.Sprint(new)
			changeKeyUnsafe("participants", newstring, "news/"+p.ID+"/")
			return c.String(http.StatusOK, "success")
		}
		return c.String(http.StatusOK, "news with id is not an event")

	}
	return c.String(http.StatusForbidden, "forbidden")

}

func removeparticipant(c echo.Context) error {
	p := new(removeNews)
	if err := c.Bind(p); err != nil {
		return err
	}
	if p.PSWD == serverpswd {
		old := readKeyUnsafe("participants", "news/"+p.ID)
		if old != "nv" {
			new, _ := strconv.Atoi(old)
			new = new - 1
			newstring := fmt.Sprint(new)
			changeKeyUnsafe("participants", newstring, "news/"+p.ID+"/")
			return c.String(http.StatusOK, "success")
		}
		return c.String(http.StatusOK, "news with id is not an event")

	}
	return c.String(http.StatusForbidden, "forbidden")

}
