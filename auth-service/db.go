package main

import (
	"database/sql"
	"log"
	"os"

	// "github.com/jmoiron/sqlx"
	"github.com/RaymondAkachi/VAULT-API/auth-service/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// var DB *sqlx.DB
var  conn *sql.DB
var queries *database.Queries
// var queries &Queries
// type apiConfig struct {
// 	DB *database.Queries{db: }
// }
// quer

func InitDB(){
	godotenv.Load()

	dbURL := os.Getenv("DB_URL")
	// DB_dsn:= os.Getenv("DB_dsn")
	// dsn := "host=localhost port=5432 user=postgres password=yourpassword dbname=authdb sslmode=disable"
	// var err error
	// DB, err = sqlx.Connect("postgres", DB_dsn)
	var err error
	conn, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
	
	// apiCfg := apiConfig{
	// 	DB: database.New(conn),
	// }
	queries = database.New(conn)
	err = conn.Ping()
	if err != nil {
		log.Fatal("Cannot ping DB:", err)
	}
	log.Println("Connected to Postgres successfully")
}