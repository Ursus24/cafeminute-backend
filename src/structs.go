package main

type addProduct struct {
	TITLE       string `json:"title" xml:"title" form:"title" query:"title"`
	ALLERGENIC  string `json:"allergenic" xml:"allergenic" form:"allergenic" query:"allergenic"`
	PRIZE       string `json:"prize" xml:"prize" form:"prize" query:"prize"`
	DESCRIPTION string `json:"description" xml:"description" form:"description" query:"description"`
	SALE        string `json:"sale" xml:"sale" form:"sale" query:"sale"`
	CALORIES    string `json:"calories" xml:"calories" form:"calories" query:"calories"`
	IMAGE       string `json:"IMAGE" xml:"IMAGE" form:"IMAGE" query:"IMAGE"`
}

type getProduct struct {
	ID string `json:"id" xml:"id" form:"id" query:"id"`
}

type addNotification struct {
	HEADING string `json:"heading" xml:"heading" form:"heading" query:"heading"`
	CONTENT string `json:"content" xml:"content" form:"content" query:"content"`
	DATE    string `json:"date" xml:"date" form:"date" query:"date"`
	TIME    string `json:"time" xml:"time" form:"time" query:"time"`
}

var daysMonth = [12]int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
