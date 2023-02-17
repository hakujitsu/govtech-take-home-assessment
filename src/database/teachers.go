package database

import (
	"assignment/teacher-api/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func AddTeacherToDB(tx *sqlx.Tx, email string) (int, error) {
	_, err := tx.Exec("INSERT IGNORE INTO teachers (email) VALUES (?)", email)
	if err != nil {
		fmt.Printf("%v", err)
		return 0, fmt.Errorf("AddTeacherToDB: %v", err)
	}

	teacherRow := tx.QueryRow("SELECT * FROM teachers WHERE email = ?", email)
	if err != nil {
		fmt.Printf("%v", err)
		return 0, fmt.Errorf("AddTeacherToDB: %v", err)
	}

	teacher, err := models.ReadRowAsTeacher(teacherRow)
	if err != nil {
		fmt.Printf("%v", err)
		return 0, fmt.Errorf("AddTeacherToDB: %v", err)
	}

	return teacher.ID, nil
}
