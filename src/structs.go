package main

var daysMonth = [12]int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

type addProduct struct {
	PSWD        string `json:"pswd" xml:"pswd" form:"pswd" query:"pswd"`
	TITLE       string `json:"title" xml:"title" form:"title" query:"title"`
	ALLERGENIC  string `json:"allergenic" xml:"allergenic" form:"allergenic" query:"allergenic"`
	PRIZE       string `json:"prize" xml:"prize" form:"prize" query:"prize"`
	DESCRIPTION string `json:"description" xml:"description" form:"description" query:"description"`
	SALE        string `json:"sale" xml:"sale" form:"sale" query:"sale"`
	CALORIES    string `json:"calories" xml:"calories" form:"calories" query:"calories"`
	IMAGE       string `json:"IMAGE" xml:"IMAGE" form:"IMAGE" query:"IMAGE"`
}

type addNotification struct {
	PSWD    string `json:"pswd" xml:"pswd" form:"pswd" query:"pswd"`
	HEADING string `json:"heading" xml:"heading" form:"heading" query:"heading"`
	CONTENT string `json:"content" xml:"content" form:"content" query:"content"`
	DATE    string `json:"date" xml:"date" form:"date" query:"date"`
	TIME    string `json:"time" xml:"time" form:"time" query:"time"`
}
type addSchedule struct {
	MON  string `json:"mon" xml:"mon" form:"mon" query:"mon"`
	TUE  string `json:"tue" xml:"tue" form:"tue" query:"tue"`
	WED  string `json:"wed" xml:"wed" form:"wed" query:"wed"`
	THU  string `json:"thu" xml:"thu" form:"thu" query:"thu"`
	FRI  string `json:"fri" xml:"fri" form:"fri" query:"fri"`
	PSWD string `json:"pswd" xml:"pswd" form:"pswd" query:"pswd"`
}
type addNews struct {
	PSWD    string `json:"pswd" xml:"pswd" form:"pswd" query:"pswd"`
	HEADING string `json:"heading" xml:"heading" form:"heading" query:"heading"`
	CONTENT string `json:"content" xml:"content" form:"content" query:"content"`
	IMAGE   string `json:"image" xml:"image" form:"image" query:"image"`
	EVENT   string `json:"event" xml:"event" form:"event" query:"event"`
}

type getProduct struct {
	ID string `json:"id" xml:"id" form:"id" query:"id"`
}

type getNews struct {
	ID string `json:"id" xml:"id" form:"id" query:"id"`
}
type removeNews struct {
	ID   string `json:"id" xml:"id" form:"id" query:"id"`
	PSWD string `json:"pswd" xml:"pswd" form:"pswd" query:"pswd"`
}
type getNotification struct {
	ID string `json:"id" xml:"id" form:"id" query:"id"`
}
type removeNotification struct {
	ID   string `json:"id" xml:"id" form:"id" query:"id"`
	PSWD string `json:"pswd" xml:"pswd" form:"pswd" query:"pswd"`
}

type changeNotifications struct {
	PSWD  string `json:"pswd" xml:"pswd" form:"pswd" query:"pswd"`
	ID    string `json:"id" xml:"id" form:"id" query:"id"`
	KEY   string `json:"key" xml:"key" form:"key" query:"key"`
	VALUE string `json:"value" xml:"value" form:"value" query:"value"`
}

type changeNews struct {
	PSWD  string `json:"pswd" xml:"pswd" form:"pswd" query:"pswd"`
	ID    string `json:"id" xml:"id" form:"id" query:"id"`
	KEY   string `json:"key" xml:"key" form:"key" query:"key"`
	VALUE string `json:"value" xml:"value" form:"value" query:"value"`
}

type changeProduct struct {
	PSWD  string `json:"pswd" xml:"pswd" form:"pswd" query:"pswd"`
	ID    string `json:"id" xml:"id" form:"id" query:"id"`
	KEY   string `json:"key" xml:"key" form:"key" query:"key"`
	VALUE string `json:"value" xml:"value" form:"value" query:"value"`
}
type removeProduct struct {
	PSWD string `json:"pswd" xml:"pswd" form:"pswd" query:"pswd"`
	ID   string `json:"id" xml:"id" form:"id" query:"id"`
}

type setOpen struct {
	PSWD string `json:"pswd" xml:"pswd" form:"pswd" query:"pswd"`
}

type customers struct {
	PSWD      string `json:"pswd" xml:"pswd" form:"pswd" query:"pswd"`
	CUSTOMERS string `json:"customers" xml:"customers" form:"customers" query:"customers"`
}

type setClosed struct {
	PSWD string `json:"pswd" xml:"pswd" form:"pswd" query:"pswd"`
}
