package controllers

import (
	"assignment/teacher-api/services"
	"assignment/teacher-api/util"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type RetrieveForNotificationsRequest struct {
	Teacher      string `json:"teacher" validate:"required,email"`
	Notification string `json:"notification" validate:"required"`
}

type RetrieveForNotificationsResponse struct {
	Recipients []string `json:"recipients"`
}

func RetrieveForNotifications(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var data RetrieveForNotificationsRequest
	err := util.ParseRequest(r, &data)
	if err != nil {
		util.SendErrorResponse(w, err.Error())
		return
	}
	err = util.ValidateRequest(&data)
	if err != nil {
		util.SendErrorResponse(w, err.Error())
		return
	}
	students, err := services.RetrieveForNotificationsService(data.Teacher, data.Notification)
	if err != nil {
		util.SendInternalServerErrorResponse(w)
		return
	}

	var emails []string
	for _, s := range students {
		emails = append(emails, s.Email)
	}

	res := RetrieveForNotificationsResponse{
		Recipients: emails,
	}

	util.SendResponse(w, http.StatusOK, res)
}
