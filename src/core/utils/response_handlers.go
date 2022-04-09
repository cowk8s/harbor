package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// ResponseHandler provides utility to handle http response.
type ResponseHandler interface {
	Handle(*http.Response) error
}

// StatusResHandler handles the response to check if the status is expected, if not returns an error.
type StatusResHandler struct {
	status int
}

// Handle ...
func (s StatusResHandler) Handle(resp *http.Response) error {
	defer resp.Body.Close()
	if resp.StatusCode != s.status {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("unexpected status code: %d, text: %s", resp.StatusCode, string(b))
	}
	return nil
}