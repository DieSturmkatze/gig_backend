package db

import (
	"database/sql"
	"log"

	_ "github.com/glebarez/go-sqlite"
)

var driver *sql.DB
var err error

func Init() {
	driver, err = sql.Open("sqlite", "./gig.db")
	if err != nil {
		log.Fatal(err)
	}
}

func UnInit() {
	driver.Close()
}
