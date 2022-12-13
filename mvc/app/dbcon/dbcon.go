package dbcon

import (
	"database/sql"
	"fmt"
	"os"
	env "app/package/config"
)

var DBcon *sql.DB

func DBConnectAccessAll() (bool) {
	db, err := sql.Open(
		"mysql", env.DBHostname + ":@" + env.DBconnection + "(" + env.DBport + ")/" + env.DBtable)
	if err != nil {
		fmt.Println("Error while connecting to database")
		os.Exit(1)
		return false
	}

	DBcon = db
	return true
}