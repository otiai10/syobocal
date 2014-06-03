package controllers

import "github.com/revel/revel"

type Doc struct {
	*revel.Controller
}

func (c Doc) Index() revel.Result {
	return c.Render()
}
