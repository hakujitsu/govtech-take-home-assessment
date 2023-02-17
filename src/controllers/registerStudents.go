package controllers

import (
	"assignment/teacher-api/services"
	"assignment/teacher-api/util"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// TODO: validate student emails
type RegisterStudentsRequest struct {
	TeacherEmail  string   `json:"teacher" validate:"required,email"`
	StudentEmails []string `json:"students" validate:"required,dive,required,email"`
}

func RegisterStudents(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var data RegisterStudentsRequest
	err := util.ParseRequest(r, &data)
	if err != nil {
		util.SendErrorResponse(w, err.Error())
		return
	}
	err = services.RegisterStudentsToTeacherService(data.StudentEmails, data.TeacherEmail)
	if err != nil {
		util.SendInternalServerErrorResponse(w)
		return
	}

	util.SendResponse(w, http.StatusNoContent, "")
}
