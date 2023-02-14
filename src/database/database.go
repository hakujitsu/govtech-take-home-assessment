package database

import (
	"assignment/teacher-api/models"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("%v", err)
		log.Fatalf("Error loading .env file")
	}

	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "edusystem",
	}
	// Get a database handle.
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected to MySQL DB!")
}

func AddTeacherToDB(email string) (int64, error) {
	result, err := db.Exec("INSERT INTO teachers (email) VALUES (?)", email)
	if err != nil {
		fmt.Printf("%v", err)
		return 0, fmt.Errorf("AddTeacherToDB: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("%v", err)
		return 0, fmt.Errorf("AddTeacherToDB: %v", err)
	}
	return id, nil
}

func ReadTeachersFromDB() ([]models.Teacher, error) {
	results, err := db.Query("SELECT * FROM teachers;")
	if err != nil {
		fmt.Printf("%v", err)
		return nil, fmt.Errorf("ReadTeachersFromDB: %v", err)
	}

	teachers, err := models.ReadRowsAsTeacher(results)
	return teachers, err
}

func DeleteTeacherFromDB(email string) error {
	_, err := db.Exec("DELETE FROM teachers WHERE email = ?", email)
	if err != nil {
		fmt.Printf("%v", err)
		return fmt.Errorf("DeleteTeacherFromDB: %v", err)
	}
	return nil
}
