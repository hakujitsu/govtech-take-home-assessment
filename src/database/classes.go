package database

import (
	"assignment/teacher-api/models"
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"
)

// TODO: wrap in transaction
func RegisterStudentsToTeacherInDB(studentEmails []string, teacherEmail string) error {
	teacher, err := ReadTeacherFromDBWithEmail(teacherEmail)
	if err != nil {
		return err
	}

	query, args, err := sqlx.In("SELECT "+strconv.Itoa(teacher.ID)+" AS teacher_id, ID AS student_id FROM students WHERE email IN (?);", studentEmails)
	if err != nil {
		return err
	}
	query = db.Rebind(query)
	rows, err := db.Query(query, args...)
	if err != nil {
		return err
	}

	class, err := models.ReadRowsAsClass(rows)
	if err != nil {
		return err
	}

	_, err = db.NamedExec("INSERT IGNORE INTO classes (teacher_id, student_id) VALUES (:teacher_id, :student_id)", class)
	if err != nil {
		return fmt.Errorf("RegisterStudentsToTeacherInDB: %v", err)
	}

	return nil
}
