package api

import (
	"encoding/json"
	"net/http"

	"github.com/cowk8s/harbor/src/lib/errors"
	lib_http "github.com/cowk8s/harbor/src/lib/http"
	"github.com/docker/distribution/registry/storage/driver"
)

// HandleInternalServerError ...
func HandleInternalServerError(w http.ResponseWriter, err error) {
	HandleError(w, errors.UnknownError(err))
}

// HandleNotMethodAllowed ...
func HandleNotMethodAllowed(w http.ResponseWriter) {
	HandleError(w, errors.MethodNotAllowedError(nil))
}

// HandleBadRequest ...
func HandleBadRequest(w http.ResponseWriter, err error) {
	HandleError(w, errors.BadRequestError(err))
}

// HandleError ...
func HandleError(w http.ResponseWriter, err error) {
	if _, ok := err.(driver.PathNotFoundError); ok {
		err = errors.New(nil).WithCode(errors.NotFoundCode).WithMessage(err.Error())
	}
	lib_http.SendError(w, err)
}

// WriteJSON response status code will be written automatically if there is an error
func WriteJSON(w http.ResponseWriter, v interface{}) error {
	b, err := json.Marshal(v)
	if err != nil {
		HandleInternalServerError(w, err)
		return err
	}

	if _, err = w.Write(b); err != nil {
		return err
	}
	return nil
}
