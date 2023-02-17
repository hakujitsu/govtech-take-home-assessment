package models

import (
	"database/sql"
	"fmt"
)

type Teacher struct {
	ID    int    `db:"id"`
	Email string `db:"email"`
}

func ReadRowAsTeacher(result *sql.Row) (Teacher, error) {
	var t Teacher
	err := result.Scan(&t.ID, &t.Email)
	if err != nil {
		return Teacher{}, fmt.Errorf("ReadRowAsTeacher: %v", err)
	}

	return t, nil
}

func ReadRowsAsTeacher(results *sql.Rows) ([]Teacher, error) {
	var teachers = make([]Teacher, 0)
	for results.Next() {
		var t Teacher
		err := results.Scan(&t.ID, &t.Email)
		if err != nil {
			return nil, fmt.Errorf("ReadRowsAsTeacher: %v", err)
		}
		teachers = append(teachers, t)
	}

	return teachers, nil
}
