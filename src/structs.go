package main

import "gorm.io/gorm"

type getMSG struct {
	SERVERPSWD string `json:"spswd" xml:"spswd" form:"spswd" query:"spswd"`
	Name       string `json:"by" xml:"by" form:"by" query:"by"`
	PASSWORD   string `json:"pswd" xml:"pswd" form:"pswd" query:"pswd"`
	SENDTO     string `json:"to" xml:"to" form:"to" query:"to"`
}

type Film struct {
	gorm.Model
	ID   uint `gorm:"primaryKey"`
	Name string
	IMDB int
}
