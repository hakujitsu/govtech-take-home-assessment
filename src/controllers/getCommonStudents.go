package controllers

import (
	"assignment/teacher-api/services"
	"assignment/teacher-api/util"
	"net/http"
	"sort"

	"github.com/julienschmidt/httprouter"
)

type GetCommonStudentsRequest struct {
	Teachers []string `json:"teachers" validate:"required,min=1,dive,required,email"`
}

type GetCommonStudentsResponse struct {
	Students []string `json:"students"`
}

func GetCommonStudents(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	queryValues := r.URL.Query()
	teachers := queryValues["teacher"]

	// Ensure all teachers are unique
	sort.Strings(teachers)
	var filteredTeachers []string
	for _, t := range teachers {
		if len(filteredTeachers) == 0 || filteredTeachers[len(filteredTeachers)-1] != t {
			filteredTeachers = append(filteredTeachers, t)
		}
	}

	data := GetCommonStudentsRequest{
		Teachers: teachers,
	}
	err := util.ValidateRequest(&data)
	if err != nil {
		util.SendErrorResponse(w, err.Error())
		return
	}

	students, err := services.GetCommonStudentsService(filteredTeachers)
	if err != nil {
		util.SendInternalServerErrorResponse(w)
		return
	}

	var emails []string
	for _, s := range students {
		emails = append(emails, s.Email)
	}

	res := GetCommonStudentsResponse{
		Students: emails,
	}

	util.SendResponse(w, http.StatusOK, res)
}
