package teachers

import (
	"net/http"

	"assignment/teacher-api/models"
	"assignment/teacher-api/services"
	"assignment/teacher-api/util"

	"github.com/julienschmidt/httprouter"
)

type CreateTeacherRequest struct {
	Email string `json:"email"`
}

type CreateTeacherResponse struct {
	Teacher models.Teacher `json:"teacher"`
}

func CreateTeacher(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var data CreateTeacherRequest
	err := util.ParseRequest(r, &data)
	if err != nil {
		util.SendErrorResponse(w, err.Error())
		return
	}

	teacher, err := services.CreateTeacherService(data.Email)

	if err != nil {
		util.SendInternalServerErrorResponse(w)
		return
	}

	res := CreateTeacherResponse{
		Teacher: teacher,
	}

	util.SendResponse(w, http.StatusCreated, res)
}
