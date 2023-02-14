package teachers

import (
	"assignment/teacher-api/services"
	"assignment/teacher-api/util"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func DeleteTeacher(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	queryValues := r.URL.Query()
	email := queryValues.Get("email")
	err := services.DeleteTeacherService(email)

	if err != nil {
		util.SendInternalServerErrorResponse(w)
		return
	}

	util.SendResponse(w, http.StatusNoContent, "")
}
