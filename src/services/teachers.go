package services

import (
	"assignment/teacher-api/database"
	"assignment/teacher-api/models"
	"fmt"
)

func CreateTeacherService(email string) (int64, error) {
	id, err := database.AddTeacherToDB(email)

	if err != nil {
		return 0, fmt.Errorf("CreateTeacher: %v", err)
	}
	return id, nil
}

func DeleteTeacherService(email string) error {
	err := database.DeleteTeacherFromDB(email)

	if err != nil {
		return fmt.Errorf("DeleteTeacher: %v", err)
	}
	return nil
}

func GetTeachersService() ([]models.Teacher, error) {
	teachers, err := database.ReadTeachersFromDB()

	if err != nil {
		return nil, fmt.Errorf("CreateTeacher: %v", err)
	}
	return teachers, nil
}
