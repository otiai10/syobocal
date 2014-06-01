package main

import "os"
import "flag"
import "fmt"
import "github.com/agtorre/gocolorize"

type Command interface {
	Name() string
	Run()
}

var commands = []Command{
	initHelp(),
	initCrawl(),
}
var (
	yellow = gocolorize.NewColor("yellow")
	green  = gocolorize.NewColor("green")
)

func main() {
	cmd := rescueCommand()
	cmd.Run()
}
func rescueCommand() (cmd Command) {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		return initHelp()
	}
	return searchCommand(args[0])
}
func searchCommand(name string) (cmd Command) {
	for _, cmd = range commands {
		if cmd.Name() == name {
			return cmd
		}
	}
	err("Command `%s` not found.", name)
	return
}
func err(format string, args ...interface{}) {
	word := fmt.Sprintf(format, args...) + "\n"
	fmt.Fprintf(os.Stderr, yellow.Paint(word))
	os.Exit(1)
}
