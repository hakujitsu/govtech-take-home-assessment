package students

import (
	"net/http"

	"assignment/teacher-api/models"
	"assignment/teacher-api/services"
	"assignment/teacher-api/util"

	"github.com/julienschmidt/httprouter"
)

type CreateStudentRequest struct {
	Email       string `json:"email"`
	IsSuspended bool   `json:"is_suspended"`
}

type CreateStudentResponse struct {
	Student models.Student `json:"student"`
}

func CreateStudent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var data CreateStudentRequest
	err := util.ParseRequest(r, &data)
	if err != nil {
		util.SendErrorResponse(w, err.Error())
		return
	}

	student, err := services.CreateStudentService(data.Email, data.IsSuspended)

	if err != nil {
		util.SendInternalServerErrorResponse(w)
		return
	}

	res := CreateStudentResponse{
		Student: student,
	}

	util.SendResponse(w, http.StatusCreated, res)
}
