package repo

import "animapi/infrastructure/db"
import "database/sql"

var Repo interface{}

type RepoClient struct {
	Db *infra.Db
}

func (client RepoClient) FindOne(dsn, id string) *sql.Row {
	return client.Db.FindOne(dsn, id)
}
