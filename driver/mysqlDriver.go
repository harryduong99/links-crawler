package driver

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Mysql struct {
	Client sql.DB
}

var MysqlDB = &Mysql{}

func (mysql *Mysql) ConnectDatabase() {
	db, dbErr := sql.Open("mysql", os.Getenv("MYSQL_CONNECTION"))
	if dbErr != nil {
		panic(dbErr.Error())
	}
	MysqlDB.Client = *db
	// defer the close till after the main function has finished
	// executing
	defer db.Close()
}
