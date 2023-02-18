package controllers

import (
	"assignment/teacher-api/util"
	"testing"
)

func TestValidateValidStudentsRequest_Valid(t *testing.T) {
	req := RegisterStudentsRequest{
		TeacherEmail:  "t_tan@gmail.com",
		StudentEmails: []string{"ann@gmail.com", "ben@gmail.com"},
	}
	err := util.ValidateRequest(&req)
	if err != nil {
		t.FailNow()
	}
}

func TestValidateRegisterStudentsRequest_MissingTeacherEmail(t *testing.T) {
	req := RegisterStudentsRequest{
		StudentEmails: []string{"ann@gmail.com", "ben@gmail.com"},
	}
	err := util.ValidateRequest(&req)
	if err == nil {
		t.FailNow()
	}
}

func TestValidateRegisterStudentsRequest_InvalidTeacherEmail(t *testing.T) {
	req := RegisterStudentsRequest{
		TeacherEmail:  "not an email",
		StudentEmails: []string{"ann@gmail.com", "ben@gmail.com"},
	}
	err := util.ValidateRequest(&req)
	if err == nil {
		t.FailNow()
	}
}

func TestValidateRegisterStudentsRequest_MissingStudentEmails(t *testing.T) {
	req := RegisterStudentsRequest{
		TeacherEmail: "t_tan@gmail.com",
	}
	err := util.ValidateRequest(&req)
	if err == nil {
		t.FailNow()
	}
}

func TestValidateRegisterStudentsRequest_EmptyStudentEmails(t *testing.T) {
	req := RegisterStudentsRequest{
		TeacherEmail:  "t_tan@gmail.com",
		StudentEmails: []string{},
	}
	err := util.ValidateRequest(&req)
	if err == nil {
		t.FailNow()
	}
}

func TestValidateRegisterStudentsRequest_InvalidStudentEmails(t *testing.T) {
	req := RegisterStudentsRequest{
		TeacherEmail:  "t_tan@gmail.com",
		StudentEmails: []string{"anngmail.com", "ben@gmail.com"},
	}
	err := util.ValidateRequest(&req)
	if err == nil {
		t.FailNow()
	}
}

func TestValidateGetCommonStudentsRequest_Valid(t *testing.T) {
	req := GetCommonStudentsRequest{
		Teachers: []string{"t_tan@gmail.com", "t_teo@gmail.com"},
	}
	err := util.ValidateRequest(&req)
	if err != nil {
		t.FailNow()
	}
}

func TestValidateGetCommonStudentsRequest_MissingTeacherEmails(t *testing.T) {
	req := GetCommonStudentsRequest{}
	err := util.ValidateRequest(&req)
	if err == nil {
		t.FailNow()
	}
}

func TestValidateGetCommonStudentsRequest_EmptyTeacherEmailsSlice(t *testing.T) {
	req := GetCommonStudentsRequest{
		Teachers: []string{},
	}
	err := util.ValidateRequest(&req)
	if err == nil {
		t.FailNow()
	}
}

func TestValidateGetCommonStudentsRequest_InvalidTeacherEmail(t *testing.T) {
	req := GetCommonStudentsRequest{
		Teachers: []string{"t_tangmail.com", "t_teo@gmail.com"},
	}
	err := util.ValidateRequest(&req)
	if err == nil {
		t.FailNow()
	}
}

func TestValidateSuspendStudentRequest_Valid(t *testing.T) {
	req := SuspendStudentRequest{
		Email: "ann@gmail.com",
	}
	err := util.ValidateRequest(&req)
	if err != nil {
		t.FailNow()
	}
}

func TestValidateSuspendStudentRequest_MissingStudentEmail(t *testing.T) {
	req := SuspendStudentRequest{}
	err := util.ValidateRequest(&req)
	if err == nil {
		t.FailNow()
	}
}

func TestValidateSuspendStudentRequest_InvalidStudentEmail(t *testing.T) {
	req := SuspendStudentRequest{
		Email: "google.com",
	}
	err := util.ValidateRequest(&req)
	if err == nil {
		t.FailNow()
	}
}

func TestValidateRetrieveForNotificationsRequest_Valid(t *testing.T) {
	req := RetrieveForNotificationsRequest{
		Teacher:      "t_tan@gmail.com",
		Notification: "Hello",
	}
	err := util.ValidateRequest(&req)
	if err != nil {
		t.FailNow()
	}
}

func TestValidateRetrieveForNotificationsRequest_MissingTeacher(t *testing.T) {
	req := RetrieveForNotificationsRequest{
		Notification: "Hello",
	}
	err := util.ValidateRequest(&req)
	if err == nil {
		t.FailNow()
	}
}

func TestValidateRetrieveForNotificationsRequest_InvalidTeacher(t *testing.T) {
	req := RetrieveForNotificationsRequest{
		Teacher:      "google.com",
		Notification: "Hello",
	}
	err := util.ValidateRequest(&req)
	if err == nil {
		t.FailNow()
	}
}

func TestValidateRetrieveForNotificationsRequest_MissingNotification(t *testing.T) {
	req := RetrieveForNotificationsRequest{
		Teacher: "t_tan@gmail.com",
	}
	err := util.ValidateRequest(&req)
	if err == nil {
		t.FailNow()
	}
}
