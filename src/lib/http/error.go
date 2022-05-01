package http

import (
	"fmt"
	"net/http"
	"strings"

	commonhttp "github.com/cowk8s/harbor/src/common/http"
	"github.com/cowk8s/harbor/src/lib/errors"
	openapi "github.com/go-openapi/errors"
)

func SendError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	statusCode, errPayload, _ := apiError(err)
	// the error detail is logged only, and will not be sent to client to avoid leaking server information
	if statusCode >= http.StatusInternalServerError {
		err = errors.New(nil).WithCode(errors.GeneralCode).WithMessage("internal server error")
		errPayload = errors.NewErrs(err).Error()
	} else {

	}
	w.WriteHeader(statusCode)
	fmt.Fprintln(w, errPayload)
}

func apiError(err error) (statusCode int, errPayload, stackTrace string) {
	code := 0
	var openAPIErr openapi.Error
	if errors.As(err, &openAPIErr) {
		// Before executing operation handler, go-swagger will bind a parameters object to a request and validate the request.
		// it will return directory when bind and validate failed.
		// The response format of the default ServerError implementation does not match the internal error response format.
		// So we needed to convert the format to the internal error response format.
		code = int(openAPIErr.Code())
		errCode := strings.Replace(strings.ToUpper(http.StatusText(code)), " ", "_", -1)
		err = errors.New(nil).WithCode(errCode).WithMessage(openAPIErr.Error())
	} else if _, ok := err.(*commonhttp.Error); ok {
		// make sure the legacy error format is align with the new one
	} else {

	}
	if code == 0 {
		code = http.StatusInternalServerError
	}
	fullStack := ""
	if _, ok := err.(*errors.Error); ok {
		fullStack = err.(*errors.Error).StackTrace()
	}
	return code, errors.NewErrs(err).Error(), fullStack
}
