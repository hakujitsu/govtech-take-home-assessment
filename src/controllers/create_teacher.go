package controllers

import (
	"fmt"
	"net/http"

	"assignment/teacher-api/services"
	"assignment/teacher-api/util"

	"github.com/julienschmidt/httprouter"
)

func CreateTeacher(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	email := ps.ByName("email")
	id, err := services.CreateTeacherService(email)

	if err != nil {
		util.SendInternalServerErrorResponse(w)
	}

	util.SendResponse(w, http.StatusCreated, fmt.Sprintf("Teacher with ID %v and email %v was created", id, email))
}
