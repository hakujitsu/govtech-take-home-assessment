package util

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func SendResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)

	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func SendErrorResponse(w http.ResponseWriter, errorMessage string) {
	res := ErrorResponse{Message: errorMessage}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	err := json.NewEncoder(w).Encode(res)

	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func SendInternalServerErrorResponse(w http.ResponseWriter) {
	res := ErrorResponse{Message: "Application encountered an error, please try again later."}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	err := json.NewEncoder(w).Encode(res)

	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}
