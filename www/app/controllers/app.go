package controllers

import "github.com/revel/revel"
import "github.com/otiai10/animapi"

import "fmt"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
  s, _ := animapi.Since("-24w")
  animes := animapi.DB("./my.conf", "test").FindAnimes(s)
	return c.Render(animes)
}
func (c App) Anime(tid int) revel.Result {
  anisongs := animapi.DB("./my.conf", "test").FindAnisongsByTID(tid)

  fmt.Println(anisongs)

  return c.Render(anisongs)
}
