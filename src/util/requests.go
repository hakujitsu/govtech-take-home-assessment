package util

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-playground/validator"
)

func ParseRequest(r *http.Request, data interface{}) error {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		return errors.New(INCORRECT_APPLICATION_TYPE)
	}

	dec := json.NewDecoder(r.Body)
	// dec.DisallowUnknownFields()
	err := dec.Decode(&data)
	if err != nil {
		return errors.New(COULD_NOT_PARSE_REQUEST)
	}

	return nil
}

func ValidateRequest(data interface{}) error {
	validate := validator.New()
	err := validate.Struct(data)
	if err != nil {
		// validationErrors := err.(validator.ValidationErrors)
		return errors.New(INVALID_REQUEST_SYNTAX)
	}
	return nil
}
