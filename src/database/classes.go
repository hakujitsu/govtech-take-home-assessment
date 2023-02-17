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

func GetCommonStudentsFromDB(teachers []string) ([]models.Student, error) {
	query, args, err := sqlx.In("SELECT DISTINCT students.* FROM students "+
		"INNER JOIN classes ON students.id = classes.student_id "+
		"INNER JOIN teachers ON teachers.id = classes.teacher_id "+
		"WHERE teachers.email in (?) GROUP BY students.email "+
		"HAVING COUNT(*) = "+strconv.Itoa(len(teachers))+";", teachers)
	if err != nil {
		return nil, err
	}
	query = db.Rebind(query)
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return models.ReadRowsAsStudents(rows)
}

func GetUnsuspendedStudentsFromTeacher(teacherEmail string, studentsEmails []string) ([]models.Student, error) {
	doesTeacherExist, err := db.Query("SELECT 1 FROM teachers WHERE email = ?", teacherEmail)
	if err != nil {
		return nil, err
	} else if !doesTeacherExist.Next() {
		return nil, err
	}

	query, args, err := sqlx.In("SELECT DISTINCT students.* FROM "+
		"(SELECT * FROM students WHERE students.is_suspended = FALSE) AS students "+
		"INNER JOIN classes ON students.id = classes.student_id "+
		"INNER JOIN teachers ON teachers.id = classes.teacher_id "+
		"WHERE teachers.email = ? OR students.email IN (?)", teacherEmail, studentsEmails)

	if err != nil {
		return nil, err
	}
	query = db.Rebind(query)
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	students, err := models.ReadRowsAsStudents(rows)
	if err != nil {
		return nil, err
	}

	return students, nil
}
