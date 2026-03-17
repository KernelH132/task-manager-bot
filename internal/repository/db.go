package repository

import (
	"database/sql"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" //driver to talk to postgres
)

var DB *sql.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
	connStr := os.Getenv("DB_Conn")
	connStr = strings.Trim(connStr, "\"")
	if connStr == "" {
		log.Fatal("DB_Conn environment variable is not set")
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("error opening database:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("error connecting to database:", err)
	}

	DB = db
}
