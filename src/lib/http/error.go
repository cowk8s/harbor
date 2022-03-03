package http

import (
	"fmt"
	"net/http"
	"strings"

	openapi "github.com/go-openapi/errors"

	"github.com/cowk8s/harbor/src/lib/errors"
)

var (
	codeMap = map[string]int{
		errors.BadRequestCode: http.StatusBadRequest,
		errors.DIGESTINVALID:  http.StatusBadRequest,
	}
)

func SendError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	statusCode, errPayload, _ := apiError(err)
	if statusCode >= http.StatusInternalServerError {

	}
	w.WriteHeader(statusCode)
	fmt.Fprintln(w, errPayload)
}

func apiError(err error) (statusCode int, errPayload, stackTrace string) {
	code := 0
	var openAPIErr openapi.Error
	if errors.As(err, &openAPIErr) {
		code = int(openAPIErr.Code())
		errCode := strings.Replace(strings.ToUpper(http.StatusText(code)), " ", "_", -1)
		err = errors.New(nil).WithCode(errCode).WithMessage(openAPIErr.Error())
	} else {
		code = codeMap[errors.ErrCode(err)]
	}
	if code == 0 {
		code = http.StatusInternalServerError
	}
	fullStack := ""
	return code, errors.NewErrs(err).Error(), fullStack
}
