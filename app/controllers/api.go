package controllers

import (
	"github.com/revel/revel"
	"net/url"
	"time"
	"strconv"
	"fmt"
	"github.com/araddon/dateparse"
)

type Api struct {
	*revel.Controller
}

type Hash struct {
	Hash string ` json:"hash" xml:"hash" `
}

type Params struct {
	url.Values
}

func (c Api) Index() revel.Result {
	// Show some documentation here

	var action string = c.Action
	var date string = time.Now().Format("2006")
	var title string = "Timestamp.fun"
	return c.Render(action, title, date)
}

func (c Api) TimeToDate() revel.Result {
	var paramTimestamp string

	c.Params.Bind(&paramTimestamp, "timestamp")

	data := make(map[string]interface{})

	if paramTimestamp != "" {
		i, err := strconv.ParseInt(paramTimestamp, 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		tm := time.Unix(i, 0)
		data["date"] = tm.UTC().Format(time.RFC3339)
	}

	data["href"] = "https://timestamp.fun" + "/api/v1/time"
	return c.RenderJSON(data)
}



func (c Api) DateToTime() revel.Result {
	var paramDate string

	c.Params.Bind(&paramDate, "date")

	data := make(map[string]interface{})

	if paramDate != "" {
		t, err := dateparse.ParseAny(paramDate)
		if err != nil {
			fmt.Println(err)
		}
		data["time"] = strconv.FormatInt(t.Unix(), 10)
	}
	data["href"] = "https://timestamp.fun" + "/api/v1/date"

	return c.RenderJSON(data)
}


