package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

var DB *sqlx.DB

func InitialiseDB() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    os.Getenv("DBNET"),
		Addr:   os.Getenv("DBADDR"),
		DBName: os.Getenv("DBNAME"),
	}
	// Get a database handle.
	sqlDb, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	DB = sqlx.NewDb(sqlDb, "mysql")

	pingErr := DB.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected to MySQL DB!")
}

func InitialiseTestDB() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("TEST_DBUSER"),
		Passwd: os.Getenv("TEST_DBPASS"),
		Net:    os.Getenv("TEST_DBNET"),
		Addr:   os.Getenv("TEST_DBADDR"),
		DBName: os.Getenv("TEST_DBNAME"),
	}
	// Get a database handle.
	sqlDb, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	DB = sqlx.NewDb(sqlDb, "mysql")

	pingErr := DB.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected to MySQL Test DB!")
}
