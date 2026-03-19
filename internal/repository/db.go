package repository

import (
	"database/sql"
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" //driver to talk to postgres
)

var DB *sql.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
	connStr := os.Getenv("DATABASE_URL")
	connStr = strings.Trim(connStr, "\"")

	if connStr == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("error opening database:", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		log.Fatal("error connecting to database:", err)
	}

	DB = db
	log.Println("Database connection established")
}
