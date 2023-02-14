package models

import (
	"database/sql"
	"fmt"
)

type Teacher struct {
	ID    int
	Email string
}

type Student struct {
	ID           int
	Email        string
	Is_Suspended bool
}

func ReadRowsAsTeacher(results *sql.Rows) ([]Teacher, error) {
	var teachers = make([]Teacher, 0)
	for results.Next() {
		var t Teacher
		err := results.Scan(&t.ID, &t.Email)
		if err != nil {
			return nil, fmt.Errorf("ReadTeachersFromDB: %v", err)
		}
		teachers = append(teachers, t)
	}

	return teachers, nil
}
