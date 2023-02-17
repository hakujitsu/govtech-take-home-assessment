package database

import (
	"assignment/teacher-api/models"
	"fmt"
)

func AddTeacherToDB(email string) (int, error) {
	result, err := db.Exec("INSERT INTO teachers (email) VALUES (?)", email)
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

func ReadTeacherFromDB(id int) (models.Teacher, error) {
	result := db.QueryRow("SELECT * FROM teachers WHERE ID = ?;", id)

	return models.ReadRowAsTeacher(result)
}

func ReadTeacherFromDBWithEmail(email string) (models.Teacher, error) {
	result := db.QueryRow("SELECT * FROM teachers WHERE email = ?;", email)

	return models.ReadRowAsTeacher(result)
}

func ReadTeachersFromDB() ([]models.Teacher, error) {
	results, err := db.Query("SELECT * FROM teachers;")
	if err != nil {
		fmt.Printf("%v", err)
		return nil, fmt.Errorf("ReadTeachersFromDB: %v", err)
	}

	teachers, err := models.ReadRowsAsTeacher(results)
	return teachers, err
}

func DeleteTeacherFromDB(email string) error {
	_, err := db.Exec("DELETE FROM teachers WHERE email = ?", email)
	if err != nil {
		fmt.Printf("%v", err)
		return fmt.Errorf("DeleteTeacherFromDB: %v", err)
	}
	return nil
}
