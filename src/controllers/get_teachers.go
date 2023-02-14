package controllers

import (
	"net/http"

	"assignment/teacher-api/models"
	"assignment/teacher-api/services"
	"assignment/teacher-api/util"

	"github.com/julienschmidt/httprouter"
)

type GetTeachersResponse struct {
	Teachers []models.Teacher `json:"teachers"`
}

func GetTeachers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	teachers, err := services.GetTeachersService()

	if err != nil {
		util.SendResponse(w, http.StatusBadRequest, "Teacher was not created")
	}

	res := GetTeachersResponse{
		Teachers: teachers,
	}

	util.SendResponse(w, http.StatusOK, res)
}
