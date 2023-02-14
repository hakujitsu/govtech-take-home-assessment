package teachers

import (
	"net/http"

	"assignment/teacher-api/models"
	"assignment/teacher-api/services"
	"assignment/teacher-api/util"

	"github.com/julienschmidt/httprouter"
)

type CreateTeacherResponse struct {
	Teacher models.Teacher `json:"teacher"`
}

func CreateTeacher(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	queryValues := r.URL.Query()
	email := queryValues.Get("email")
	teacher, err := services.CreateTeacherService(email)

	if err != nil {
		util.SendInternalServerErrorResponse(w)
		return
	}

	res := CreateTeacherResponse{
		Teacher: teacher,
	}

	util.SendResponse(w, http.StatusCreated, res)
}
