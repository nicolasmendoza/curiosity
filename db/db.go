package db

import (
	"database/sql"
	"fmt"
)

package rss

import (
"database/sql"
"fmt"
)

var dataSourceName = "root:password@/rsss"
var DB *sql.DB

func init() {
	// The returned DB is safe for concurrent use by multiple goroutines
	// and maintains its own pool of idle connections. Thus, the Open
	// function should be called just once. It is rarely necessary to
	// close a DB.
	fmt.Println("Initializing database!!!!")
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error)
	}
	DB = db
}

