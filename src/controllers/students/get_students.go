package students

import (
	"net/http"

	"assignment/teacher-api/models"
	"assignment/teacher-api/services"
	"assignment/teacher-api/util"

	"github.com/julienschmidt/httprouter"
)

type GetStudentsResponse struct {
	Students []models.Student `json:"students"`
}

func GetStudents(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	students, err := services.GetStudentsService()

	if err != nil {
		util.SendInternalServerErrorResponse(w)
		return
	}

	res := GetStudentsResponse{
		Students: students,
	}

	util.SendResponse(w, http.StatusOK, res)
}
