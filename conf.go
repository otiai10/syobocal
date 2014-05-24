package animapi

import "github.com/robfig/config"

type Config struct {
	cnf  *config.Config
	err  error
	opt  string
	Host string
	Port string
	User string
	Pass string
}

func File(args ...string) *Config {
	o := ""
	if 1 < len(args) {
		o = args[1]
	}
	c, e := config.ReadDefault(args[0])
	conf := &Config{cnf: c, err: e, opt: o}
	return conf.build()
}
func (c *Config) build() *Config {
	return c
}
