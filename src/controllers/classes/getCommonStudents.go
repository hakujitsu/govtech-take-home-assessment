package classes

import (
	"assignment/teacher-api/services"
	"assignment/teacher-api/util"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type GetCommonStudentsResponse struct {
	Students []string `json:"students"`
}

func GetCommonStudents(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	queryValues := r.URL.Query()
	teachers := queryValues["teacher"]

	// TODO: filter to ensure all teachers are unique
	students, err := services.GetCommonStudentsService(teachers)
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
