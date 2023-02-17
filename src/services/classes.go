package services

import (
	"assignment/teacher-api/database"
	"assignment/teacher-api/models"
	"fmt"
)

func RegisterStudentsToTeacherService(studentEmails []string, teacherEmail string) error {
	err := database.RegisterStudentsToTeacherInDB(studentEmails, teacherEmail)
	if err != nil {
		return fmt.Errorf("RegisterStudentsToTeacherService: %v", err)
	}

	return nil
}

func GetCommonStudentsService(teacherEmails []string) ([]models.Student, error) {
	students, err := database.GetCommonStudentsFromDB(teacherEmails)
	if err != nil {
		fmt.Printf("%v", err)
		return nil, fmt.Errorf("RegisterStudentsToTeacherService: %v", err)
	}

	return students, nil
}
