package main

type addProduct struct {
	title      string `json:"title" xml:"title" form:"title" query:"title"`
	allergenic string `json:"allergenic" xml:"allergenic" form:"allergenic" query:"allergenic"`
	prize      string `json:"prize" xml:"prize" form:"prize" query:"prize"`
	descrition string `json:"description" xml:"description" form:"description" query:"description"`
}
