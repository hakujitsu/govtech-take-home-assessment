package controllers

import (
	"fmt"
	"net/http"

	"assignment/teacher-api/services"
	"assignment/teacher-api/util"

	"github.com/julienschmidt/httprouter"
)

func CreateTeacher(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := services.CreateTeacherService(ps.ByName("email"))

	if err != nil {
		util.SendResponse(w, http.StatusBadRequest, "Teacher was not created")
	}

	util.SendResponse(w, http.StatusCreated, fmt.Sprintf("Teacher %v was created", id))
}
