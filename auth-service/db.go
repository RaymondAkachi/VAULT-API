package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func InitDB(){
	dsn := "host=localhost port=5432 user=postgres password=yourpassword dbname=authdb sslmode=disable"
	var err error
	DB, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)

	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Cannot ping DB:", err)
	}
	log.Println("Connected to Postgres successfully")
}