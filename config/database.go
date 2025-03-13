package config

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"os"
)

var DB *sqlx.DB

func ConnectDB() *sqlx.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Database connected successfully")

	// Run migrations
	RunMigrations(db)

	DB = db

	return db
}

//RunMigrations executes SQL files for setup

func RunMigrations(db *sqlx.DB) {
	file, err := ioutil.ReadFile("migrations/schema.sql")
	if err != nil {
		log.Fatal("Failed to read migration file:", err)
	}

	_, err = db.Exec(string(file))
	if err != nil {
		log.Fatal("Failed to execute migration:", err)
	}

	fmt.Println("Migrations applied successfully")
}
