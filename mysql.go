package animapi

import "github.com/robfig/config"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "fmt"

// import "github.com/otiai10/animapi/model"

// TODO: このファイルの一部の操作はinfrastructureに分割すべき
// import "github.com/otiai10/animapi/infrastructure"

var (
	Database = "mysql"
	Protocol = "tcp"
	Adress   = "%s:%s@%s(%s:%s)/%s"
)

type MySQL struct {
	conf conf
	Err  error
	db   *sql.DB
}
type conf struct {
	Port string
	Host string
	User string
	Pass string
	Dsn  string
}

func DB(args ...string) *MySQL {
	client := &MySQL{}
	c, e := ensureConf(args)
	if e != nil {
		client.Err = e
		return client
	}
	client.conf = c
	client.connect()
	return client
}
func (client *MySQL) connect() {
	address := fmt.Sprintf(
		Adress,
		client.conf.User,
		client.conf.Pass,
		Protocol,
		client.conf.Host,
		client.conf.Port,
		client.conf.Dsn, // TODO: ここ指定できるようにする
	)
	db, e := sql.Open(
		Database,
		address,
	)
	client.db = db
	client.Err = e
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
	ds, e := cnf.String(confOpt, "dsn")
	if e != nil {
		return
	}
	c.Dsn = ds
	return
}
