package services

import (
	"assignment/teacher-api/database"
	"assignment/teacher-api/models"
	"fmt"
	"strings"
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

// TODO: need to ensure no duplicates
func RetrieveForNotificationsService(teacher string, notification string) ([]models.Student, error) {
	mentionedStudents := parseForMentions(notification)
	students, err := database.GetUnsuspendedStudentsFromTeacher(teacher, mentionedStudents)
	if err != nil {
		fmt.Printf("%v", err)
		return nil, fmt.Errorf("RetrieveForNotificationsService: %v", err)
	}

	return students, nil
}

func parseForMentions(notification string) []string {
	splitString := strings.Fields(notification)
	var students []string
	for _, s := range splitString {
		if s[0:1] == "@" && len(s) > 1 {
			students = append(students, s[1:])
		}
	}

	return students
}
