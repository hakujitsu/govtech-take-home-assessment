package models

import (
	"database/sql"
	"fmt"
)

type Student struct {
	ID           int    `db:"id"`
	Email        string `db:"email"`
	Is_Suspended bool   `db:"is_suspended"`
}

func CreateNewStudentFromEmail(email string) Student {
	return Student{
		Email:        email,
		Is_Suspended: false,
	}
}

func CreateNewStudentsFromEmail(emails []string) []Student {
	var students []Student
	for _, e := range emails {

		students = append(students, CreateNewStudentFromEmail(e))
	}
	return students
}

func ReadRowAsStudent(result *sql.Row) (Student, error) {
	var s Student
	err := result.Scan(&s.ID, &s.Email, &s.Is_Suspended)
	if err != nil {
		return Student{}, fmt.Errorf("ReadRowAsStudent: %v", err)
	}

	return s, nil
}

func ReadRowsAsStudents(results *sql.Rows) ([]Student, error) {
	var students = make([]Student, 0)
	for results.Next() {
		var s Student
		err := results.Scan(&s.ID, &s.Email, &s.Is_Suspended)
		if err != nil {
			return nil, fmt.Errorf("ReadRowsAsStudents: %v", err)
		}
		students = append(students, s)
	}

	return students, nil
}
