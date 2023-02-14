package students

import (
	"net/http"
	"strconv"

	"assignment/teacher-api/models"
	"assignment/teacher-api/services"
	"assignment/teacher-api/util"

	"github.com/julienschmidt/httprouter"
)

type CreateStudentResponse struct {
	Student models.Student `json:"student"`
}

func CreateStudent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	queryValues := r.URL.Query()
	email := queryValues.Get("email")
	isSuspended, err := strconv.ParseBool(queryValues.Get("isSuspended"))

	if err != nil {
		// TODO: change to invalid value error when doing data validation
		util.SendInternalServerErrorResponse(w)
		return
	}

	student, err := services.CreateStudentService(email, isSuspended)

	if err != nil {
		util.SendInternalServerErrorResponse(w)
		return
	}

	res := CreateStudentResponse{
		Student: student,
	}

	util.SendResponse(w, http.StatusCreated, res)
}
