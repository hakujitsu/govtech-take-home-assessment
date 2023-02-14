package controllers

import (
	"assignment/teacher-api/services"
	"assignment/teacher-api/util"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func DeleteTeacher(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := services.DeleteTeacherService(ps.ByName("email"))

	if err != nil {
		util.SendResponse(w, http.StatusBadRequest, "Teacher was not deleted")
	}

	util.SendResponse(w, http.StatusNoContent, "Teacher was deleted")
}
