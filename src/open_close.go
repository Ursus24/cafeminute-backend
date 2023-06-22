package main

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

var open = "false"

func isopen(c echo.Context) error {
	return c.String(http.StatusOK, open)
}

func setopen(c echo.Context) error {
	p := new(setOpen)
	if err := c.Bind(p); err != nil {
		return err
	}
	if string(p.PSWD) == serverpswd {
		open = "true"
		return c.String(http.StatusOK, "sucess")
	}
	return c.String(http.StatusForbidden, "forbidden")
}

func setclosed(c echo.Context) error {
	p := new(setOpen)
	if err := c.Bind(p); err != nil {
		return err
	}
	if string(p.PSWD) == serverpswd {
		open = "false"
		return c.String(http.StatusOK, "success")
	}
	return c.String(http.StatusForbidden, "forbidden")
}
func add_schedule(c echo.Context) error {
	createDir("schedule")
	p := new(addSchedule)
	if err := c.Bind(p); err != nil {
		return err
	}
	//check if all dates are valid

	if p.FRI == "" || p.MON == "" || p.TUE == "" || p.WED == "" || p.THU == "" {
		return c.String(http.StatusBadRequest, "bad request")
	}
	//join all values to array
	var openingtimes = []string{p.MON, p.TUE, p.WED, p.THU, p.FRI}
	if string(p.PSWD) == serverpswd {
		if !hasKey("schedule", "schedule") {
			addKeyUnsafe("schedule", strings.Join(openingtimes, "◌◌◞◌◌◌"), "schedule")
		} else {
			changeKeyUnsafe("schedule", strings.Join(openingtimes, "◌◌◞◌◌◌"), "schedule")
		}
		return c.String(http.StatusOK, "success")
	}
	return c.String(http.StatusForbidden, "forbidden")
}

func get_schedule(c echo.Context) error {
	if !hasKey("schedule", "schedule") {
		return c.String(http.StatusConflict, "no schedule")
	}
	return c.String(http.StatusOK, readKeyUnsafe("schedule", "schedule"))
}
