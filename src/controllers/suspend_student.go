package controllers

import (
	"assignment/teacher-api/services"
	"assignment/teacher-api/util"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type SuspendStudentRequest struct {
	Email string `json:"student" validate:"required,email"`
}

func SuspendStudent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var data SuspendStudentRequest
	err := util.ParseRequest(r, &data)
	if err != nil {
		util.SendErrorResponse(w, err.Error())
		return
	}
	err = util.ValidateRequest(&data)
	if err != nil {
		util.SendErrorResponse(w, err.Error())
		return
	}

	err = services.SuspendStudentService(data.Email, true)

	if err != nil {
		util.SendErrorResponse(w, err.Error())
		return
	}

	util.SendResponse(w, http.StatusNoContent, "")
}
