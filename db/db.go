package db

import (
	"database/sql"
	"fmt"
)

var Pool *sql.DB // pool connection...

const (
	schema = "root:password@/rsss" // Schema name
	driver = "mysql"               // Driver name
)

func Get() {
	if Pool == nil {
		fmt.Println("Preparing Database connection...")
		initConnection()
	}

}

// The returned DB is safe for concurrent use by multiple goroutines
// and maintains its own pool of idle connections. Thus, the Open
// function should be called just once. It is rarely necessary to
// close a DB.
func initConnection() {
	db, err := sql.Open(driver, schema)
	if err != nil {
		panic(err.Error())
	}
	// set to db to Pool
	Pool = db
}
