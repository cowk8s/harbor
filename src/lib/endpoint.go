package lib

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/cowk8s/harbor/src/lib/errors"
)

// ValidateHTTPURL checks whether the provided string is a valid HTTP URL.
// If it is, return the URL in format "scheme://host:port" to avoid the SSRF
func ValidateHTTPURL(s string) (string, error) {
	s = strings.Trim(s, " ")
	s = strings.TrimRight(s, "/")
	if len(s) == 0 {
		return "", errors.New(nil).WithCode(errors.BadRequestCode).WithMessage("empty string")
	}
	if !strings.Contains(s, "://") {
		s = "http://" + s
	}
	url, err := url.Parse(s)
	if err != nil {
		return "", errors.New(nil).WithCode(errors.BadRequestCode).WithMessage("invalid URL: %s", err.Error())
	}
	if url.Scheme != "http" && url.Scheme != "https" {
		return "", errors.New(nil).WithCode(errors.BadRequestCode).WithMessage("invalid HTTP scheme: %s", url.Scheme)
	}
	// To avoid SSRF security issue, refer to #3755 for more detail
	return fmt.Sprintf("%s://%s%s", url.Scheme, url.Host, url.Path), nil
}
