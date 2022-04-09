package lib

import "net/http"

// NewResponseRecorder creates a response recorder
func NewResponseRecorder(w http.ResponseWriter) *ResponseRecorder {
	recorder := &ResponseRecorder{}
	recorder.ResponseWriter = w
	return recorder
}

// ResponseRecorder is a wrapper for the http.ResponseWriter to record the response status code
type ResponseRecorder struct {
	StatusCode  int
	wroteHeader bool
	http.ResponseWriter
}

// Write records the status code before writing data to the underlying writer
func (r *ResponseRecorder) Write(data []byte) (int, error) {
	if !r.wroteHeader {
		r.WriteHeader(http.StatusOK)
	}
	return r.ResponseWriter.Write(data)
}

// WriteHeader records the status code before writing the code to the underlying writer
func (r *ResponseRecorder) WriteHeader(statusCode int) {
	if !r.wroteHeader {
		r.wroteHeader = true
		r.StatusCode = statusCode
		r.ResponseWriter.WriteHeader(statusCode)
	}
}

// Success checks whether the status code is >= 200 & <= 399
func (r *ResponseRecorder) Success() bool {
	statusCode := r.StatusCode
	if statusCode == 0 {
		// NOTE: r.code is zero means that `WriteHeader` not called by the http handler,
		// so process it as http.StatusOK
		statusCode = http.StatusOK
	}

	return statusCode >= http.StatusOK && statusCode < http.StatusBadRequest
}
