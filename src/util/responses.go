package util

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SendResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)

	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func SendInternalServerErrorResponse(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	err := json.NewEncoder(w).Encode("Application encountered an error, please try again later.")

	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}
