package services

import (
	"assignment/teacher-api/database"
	"fmt"
)

func RegisterStudentsToTeacherService(studentEmails []string, teacherEmail string) error {
	err := database.RegisterStudentsToTeacherInDB(studentEmails, teacherEmail)
	if err != nil {
		return fmt.Errorf("RegisterStudentsToTeacherService: %v", err)
	}

	return nil
}
