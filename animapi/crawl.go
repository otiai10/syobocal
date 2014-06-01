package main

import "github.com/otiai10/animapi"
import "github.com/otiai10/animapi/model"

type CmdCrawl struct {
	count int
}

func initCrawl() *CmdCrawl {
	return &CmdCrawl{
		count: 0,
	}
}
func (c *CmdCrawl) Name() string {
	return "crawl"
}
func (c *CmdCrawl) Run() {
}
func (c *CmdCrawl) execute() (e error) {
	c.count++
	programs, e := animapi.SYOBOCAL.FindProgramsSince("-1w")
	if e != nil {
		return
	}
	return c.store(programs)
}
func (c *CmdCrawl) store(programs []model.Program) (e error) {
	db := animapi.DB("./my.conf", "test")
	if e = db.AddPrograms(programs); e != nil {
		return
	}
	for _, program := range programs {
		if e = db.AddAnime(program.Anime); e != nil {
			return
		}
		if e = db.AddAnisongsOfAnime(program.Anime); e != nil {
			return
		}
	}
	return
}
