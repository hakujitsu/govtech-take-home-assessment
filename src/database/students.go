package database

import (
	"assignment/teacher-api/models"
	"fmt"
)

func AddStudentToDB(email string, isSuspended bool) (int, error) {
	result, err := db.Exec("INSERT INTO students (email, is_suspended) VALUES (?, ?)", email, isSuspended)
	if err != nil {
		fmt.Printf("%v", err)
		return 0, fmt.Errorf("AddStudentToDB: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("%v", err)
		return 0, fmt.Errorf("AddStudentToDB: %v", err)
	}
	return int(id), nil
}

func ReadStudentFromDB(id int) (models.Student, error) {
	result := db.QueryRow("SELECT * FROM students WHERE ID = ?;", id)

	return models.ReadRowAsStudent(result)
}

func ReadStudentsFromDB() ([]models.Student, error) {
	results, err := db.Query("SELECT * FROM students;")
	if err != nil {
		fmt.Printf("%v", err)
		return nil, fmt.Errorf("ReadStudentsFromDB: %v", err)
	}

	students, err := models.ReadRowsAsStudents(results)
	return students, err
}

func DeleteStudentFromDB(email string) error {
	_, err := db.Exec("DELETE FROM students WHERE email = ?", email)
	if err != nil {
		fmt.Printf("%v", err)
		return fmt.Errorf("DeleteStudentFromDB: %v", err)
	}
	return nil
}
