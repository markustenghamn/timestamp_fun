package controllers

import (
	"github.com/revel/revel"
	"time"
	"strconv"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	var date string = time.Now().Format("2006")
	var title string = "Timestamp.fun"
	var count string = "5"

	t := time.Now()
	timestampVar := strconv.FormatInt(t.Unix(), 10)

	dateVar := time.Now().UTC().Format(time.RFC3339)

	var stringsTempVar string
	var difficultyTempVar int
	c.Params.Query=c.Request.URL.Query()
	c.Params.Bind(&stringsTempVar,"strings")
	c.Params.Bind(&difficultyTempVar,"difficulty")

	// Should be moved to new controller and all controllers inherit
	var action string = c.Action
	return c.Render(date, title, count, timestampVar, dateVar, action)
}
