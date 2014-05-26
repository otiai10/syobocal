package animapi

import "github.com/robfig/config"

type MySqlClient struct {
	conf conf
	Err  error
}
type conf struct {
	Port string
	Host string
	User string
	Pass string
}

func DB(args ...string) (client MySqlClient) {
	c, e := ensureConf(args)
	if e != nil {
		client.Err = e
		return
	}
	client.conf = c
	return
}
func ensureConf(args []string) (c conf, e error) {
	confPath := args[0]
	confOpt := "DEFAULT"
	if 1 < len(args) {
		confOpt = args[1]
	}
	cnf, e := config.ReadDefault(confPath)
	if e != nil {
		return
	}

	po, e := cnf.String(confOpt, "port")
	if e != nil {
		return
	}
	c.Port = po
	ho, e := cnf.String(confOpt, "host")
	if e != nil {
		return
	}
	c.Host = ho
	us, e := cnf.String(confOpt, "user")
	if e != nil {
		return
	}
	c.User = us
	pa, e := cnf.String(confOpt, "pass")
	if e != nil {
		return
	}
	c.Pass = pa
	return
}
