package animapi

import "github.com/robfig/config"

type Config struct {
	cnf  *config.Config
	Err  error
	opt  string
	Host string
	Port string
	User string
	Pass string
}

var CONF = proxyConf{}

type proxyConf struct{}

func (proxy proxyConf) File(args ...string) *Config {
	return conffile(args)
}

func conffile(args []string) *Config {
	o := ""
	if 1 < len(args) {
		o = args[1]
	}
	c, e := config.ReadDefault(args[0])
	conf := &Config{cnf: c, Err: e, opt: o}
	return conf.build()
}
func (c *Config) build() *Config {
	// TODO: judge Err and return early
	c.Host, c.Err = c.cnf.String(c.opt, "host")
	c.Port, c.Err = c.cnf.String(c.opt, "port")
	c.User, c.Err = c.cnf.String(c.opt, "user")
	c.Pass, c.Err = c.cnf.String(c.opt, "pass")
	return c
}
