package util

import (
	"encoding/json"
	"errors"
	"net/http"
)

func ParseRequest(r *http.Request, data interface{}) error {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		return errors.New("Content Type is not application/json")
	}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&data)
	if err != nil {
		// TODO: handle logging
		return errors.New("Could not parse request body")
	}

	return nil
}
