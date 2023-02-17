package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func AddTeacherToDB(tx *sqlx.Tx, email string) (int, error) {
	result, err := tx.Exec("INSERT IGNORE INTO teachers (email) VALUES (?)", email)
	if err != nil {
		fmt.Printf("%v", err)
		return 0, fmt.Errorf("AddTeacherToDB: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("%v", err)
		return 0, fmt.Errorf("AddTeacherToDB: %v", err)
	}
	return int(id), nil
}
