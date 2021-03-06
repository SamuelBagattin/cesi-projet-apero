package config

import (
	"database/sql"
	"fmt"
	"os"
	"sync"
)

var once sync.Once

type DriverPg struct {
	conn string
}

// variable globale
var db *sql.DB

func DatabaseInit() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=require",
		os.Getenv("dbHost"), os.Getenv("dbPort"), os.Getenv("dbUser"), os.Getenv("dbPassword"), os.Getenv("dbName"))
	once.Do(func() {
		var err error
		db, err = sql.Open("postgres", psqlInfo)
		if err != nil {
			panic(err)
		}
	})

	return db
}
