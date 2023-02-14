package controllers

import (
	"net/http"

	"assignment/teacher-api/models"
	"assignment/teacher-api/services"
	"assignment/teacher-api/util"

	"github.com/julienschmidt/httprouter"
)

type CreateTeachersResponse struct {
	Teacher models.Teacher `json:"teacher"`
}

func CreateTeacher(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	email := ps.ByName("email")
	teacher, err := services.CreateTeacherService(email)

	if err != nil {
		util.SendInternalServerErrorResponse(w)
	}

	res := CreateTeachersResponse{
		Teacher: teacher,
	}

	util.SendResponse(w, http.StatusCreated, res)
}
