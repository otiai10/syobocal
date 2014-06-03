package controllers

import "github.com/revel/revel"

type Docs struct {
	*revel.Controller
}

func (c Docs) Index() revel.Result {
	return c.Render()
}
