package driver

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

type Postgres struct {
	Client sql.DB
}

var PostgresDB = &Postgres{}

func (postgres *Postgres) ConnectDatabase() {
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_CONNECTION"))
	if err != nil {
		panic(err)
	}

	PostgresDB.Client = *db

	defer db.Close()
}
