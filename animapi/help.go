package main

type CmdHelp struct {
	name string
}

func initHelp() *CmdHelp {
	return &CmdHelp{
		name: "help",
	}
}
func (c *CmdHelp) Name() string {
	return c.name
}
func (c *CmdHelp) Run() {
	println("TODO: ここでヘルプ出す")
}
