package services

import (
	"assignment/teacher-api/database"
	"assignment/teacher-api/models"
	"fmt"
	"strings"
)

func RegisterStudentsToTeacherService(studentEmails []string, teacherEmail string) error {
	students := models.CreateNewStudentsFromEmail(studentEmails)

	tx, err := database.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	id, err := database.AddTeacherToDB(tx, teacherEmail)
	if err != nil {
		return err
	}

	err = database.AddStudentsToDB(tx, students)
	if err != nil {
		return err
	}

	err = database.RegisterStudentsToTeacherInDB(tx, studentEmails, id)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func GetCommonStudentsService(teacherEmails []string) ([]models.Student, error) {
	students, err := database.GetCommonStudentsFromDB(teacherEmails)
	if err != nil {
		return nil, fmt.Errorf("GetCommonStudentsService: %v", err)
	}

	return students, nil
}

func SuspendStudentService(email string, suspend bool) error {
	return database.UpdateStudentInDB(email, suspend)
}

func RetrieveForNotificationsService(teacher string, notification string) ([]models.Student, error) {
	mentionedStudents := parseForMentions(notification)
	students, err := database.GetUnsuspendedStudentsFromTeacher(teacher, mentionedStudents)
	if err != nil {
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
