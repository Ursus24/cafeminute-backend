package main

type addProduct struct {
	TITLE       string `json:"title" xml:"title" form:"title" query:"title"`
	ALLERGENIC  string `json:"allergenic" xml:"allergenic" form:"allergenic" query:"allergenic"`
	PRIZE       string `json:"prize" xml:"prize" form:"prize" query:"prize"`
	DESCRIPTION string `json:"description" xml:"description" form:"description" query:"description"`
}
