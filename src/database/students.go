package database

import (
	"assignment/teacher-api/models"
	"assignment/teacher-api/util"
	"errors"

	"github.com/jmoiron/sqlx"
)

func AddStudentsToDB(tx *sqlx.Tx, students []models.Student) error {
	_, err := tx.NamedExec("INSERT IGNORE INTO students (email, is_suspended) VALUES (:email, :is_suspended)", students)
	return err
}

func UpdateStudentInDB(email string, isSuspended bool) error {
	doesStudentExist, err := DB.Query("SELECT 1 FROM students WHERE email = ?", email)
	if err != nil {
		return err
	} else if !doesStudentExist.Next() {
		return errors.New(util.STUDENT_DOES_NOT_EXIST)
	}

	_, err = DB.Exec("UPDATE students SET is_suspended = ? WHERE email = ?", isSuspended, email)
	if err != nil {
		return err
	}
	return nil
}
