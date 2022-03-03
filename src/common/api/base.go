package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/cowk8s/harbor/src/lib/log"
)

const (
	defaultPageSize int64 = 500
	maxPageSize     int64 = 500

	// APIVersion is the current core api version
	APIVersion = "v2.0"
)

// BaseAPI wraps common methods for controllers to host API
type BaseAPI struct {
	beego.Controller
}

// Context returns the context.Context from http.Request
func (b *BaseAPI) Context() context.Context {
	return b.Ctx.Request.Context()
}

func (b *BaseAPI) GetStringFromPath(key string) string {
	return b.Ctx.Input.Param(key)
}

// GetInt64FromPath gets the param from path and returns it as int64
func (b *BaseAPI) GetInt64FromPath(key string) (int64, error) {
	value := b.Ctx.Input.Param(key)
	return strconv.ParseInt(value, 10, 64)
}

func (b *BaseAPI) ParamExistsInPath(key string) bool {
	return b.GetStringFromPath(key) != ""
}

func (b *BaseAPI) Render() error {
	return nil
}

func (b *BaseAPI) RenderError(code int, text string) {

}

func (b *BaseAPI) DecodeJSONReq(v interface{}) error {
	err := json.Unmarshal(b.Ctx.Input.CopyBody(1<<32), v)
	if err != nil {
		log.Errorf("Error while decoding the json request, error: %v, %v",
			err, string(b.Ctx.Input.CopyBody(1 << 32)[:]))
		return errors.New("invalid json request")
	}
	return nil
}

func (b *BaseAPI) Validate(v interface{}) (bool, error) {
	validator := validation.Validation{}
	isValid, err := validator.Valid(v)
	if err != nil {
		log.Errorf("failed to validate: %v", err)
		return false, err
	}

	if !isValid {
		message := ""
		for _, e := range validator.Errors {
			message += fmt.Sprintf("%s %s \n", e.Field, e.Message)
		}
		return false, errors.New(message)
	}
	return true, nil
}

func (b *BaseAPI) DecodeJSONReqAndValidate(v interface{}) (bool, error) {
	if err := b.DecodeJSONReq(v); err != nil {
		return false, err
	}
	return b.Validate(v)
}

func (b *BaseAPI) Redirect(statusCode int, resouceID string) {
	requestURI := b.Ctx.Request.RequestURI
	resourceURI := requestURI + "/" + resouceID

	b.Ctx.Redirect(statusCode, resourceURI)
}

func (b *BaseAPI) GetIDFromURL() (int64, error) {
	idStr := b.Ctx.Input.Param(":id")
	if len(idStr) == 0 {
		return 0, errors.New("invalid ID in URL")
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		return 0, errors.New("invalid ID in URL")
	}

	return id, nil
}

func (b *BaseAPI) SetPaginationHeader(total, page, pageSize int64) {
	b.Ctx.ResponseWriter.Header().Set("X-Total-Count", strconv.FormatInt(total, 10))

	link := ""

	// SetPaginationHeader set previous link
	if page > 1 && (page-1)*pageSize <= total {
		u := *(b.Ctx.Request.URL)
		q := u.Query()
		q.Set("page", strconv.FormatInt(page-1, 10))
		u.RawQuery = q.Encode()
		if len(link) != 0 {
			link += ", "
		}
		link += fmt.Sprintf("<%s>; rel=\"prev\"", u.String())
	}
}
