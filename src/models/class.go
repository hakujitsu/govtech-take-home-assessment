package models

import (
	"database/sql"
	"fmt"
)

type Class struct {
	TeacherID int `db:"teacher_id"`
	StudentID int `db:"student_id"`
}

func ReadRowsAsClass(results *sql.Rows) ([]Class, error) {
	var class = make([]Class, 0)
	for results.Next() {
		var c Class
		err := results.Scan(&c.TeacherID, &c.StudentID)
		if err != nil {
			return nil, fmt.Errorf("ReadRowsAsClass: %v", err)
		}
		class = append(class, c)
	}

	return class, nil
}
