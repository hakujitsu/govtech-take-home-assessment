package services

import (
	"assignment/teacher-api/database"
	"assignment/teacher-api/models"
	"fmt"
)

func CreateStudentService(email string, isSuspended bool) (models.Student, error) {
	id, err := database.AddStudentToDB(email, isSuspended)

	if err != nil {
		return models.Student{}, fmt.Errorf("CreateStudent: %v", err)
	}

	student, err := database.ReadStudentFromDB(id)
	if err != nil {
		return models.Student{}, fmt.Errorf("CreateStudent: %v", err)
	}

	return student, nil
}

func DeleteStudentService(email string) error {
	err := database.DeleteStudentFromDB(email)

	if err != nil {
		return fmt.Errorf("DeleteStudent: %v", err)
	}
	return nil
}

func GetStudentsService() ([]models.Student, error) {
	students, err := database.ReadStudentsFromDB()

	if err != nil {
		return nil, fmt.Errorf("GetStudents: %v", err)
	}
	return students, nil
}

func SuspendStudentService(email string, suspend bool) error {
	err := database.UpdateStudentInDB(email, suspend)

	return err
}
