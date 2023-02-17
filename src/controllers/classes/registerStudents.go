package classes

import (
	"assignment/teacher-api/services"
	"assignment/teacher-api/util"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type RegisterStudentsRequest struct {
	TeacherEmail  string   `json:"teacher"`
	StudentEmails []string `json:"students"`
}

func RegisterStudents(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var data RegisterStudentsRequest
	err := util.ParseRequest(r, &data)
	if err != nil {
		util.SendErrorResponse(w, err.Error())
		return
	}
	err = services.RegisterStudentsToTeacherService(data.StudentEmails, data.TeacherEmail)
	if err != nil {
		util.SendInternalServerErrorResponse(w)
		return
	}

	util.SendResponse(w, http.StatusNoContent, "")
}
