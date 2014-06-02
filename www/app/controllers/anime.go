package controllers

import "github.com/revel/revel"
import "github.com/otiai10/animapi"

type Anime struct {
	*revel.Controller
}

func (c Anime) Anisongs(tid int) revel.Result {
	anisongs := animapi.DB("./my.conf", "test").FindAnisongsByTID(tid)
	return c.RenderJson(anisongs)
}
