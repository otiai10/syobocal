package controllers

import "github.com/revel/revel"
import "github.com/otiai10/animapi"

type Api struct {
	*revel.Controller
}

var DB = animapi.DB("./my.conf", "test")

func (c Api) AnisongsByTid(tid int) revel.Result {
	anisongs := DB.FindAnisongsByTID(tid)
	return c.RenderJson(anisongs)
}

func (c Api) Animes() revel.Result {
	since, _ := animapi.Since("-1w")
	anisongs := DB.FindAnimes(since)
	return c.RenderJson(anisongs)
}
