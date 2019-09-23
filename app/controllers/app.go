package controllers

import (
	"github.com/revel/revel"
	"github.com/revel/revel/session"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	x := session.SESSION_VALUE_NOT_FOUND
	revel.AppLog.Info(x.Error())
	return c.Render()
}
