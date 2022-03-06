package lib

import (
	"bytes"
	"errors"
	"net/http"
)

type ResponseBuffer struct {
	w           http.ResponseWriter
	code        int
	header      http.Header
	buffer      bytes.Buffer
	wroteHeader bool
	flushed     bool
}

func NewResponseBuffer(w http.ResponseWriter) *ResponseBuffer {
	return &ResponseBuffer{
		w:      w,
		header: http.Header{},
		buffer: bytes.Buffer{},
	}
}

func (r *ResponseBuffer) WriteHeader(statusCode int) {
	if r.wroteHeader {
		return
	}
	r.wroteHeader = true
	r.code = statusCode
}

func (r *ResponseBuffer) Write(data []byte) (int, error) {
	if !r.wroteHeader {
		r.WriteHeader(http.StatusOK)
	}
	return r.buffer.Write(data)
}

func (r *ResponseBuffer) Header() http.Header {
	return r.header
}

func (r *ResponseBuffer) Flush() (int, error) {
	r.flushed = true

	header := r.w.Header()
	for k, vs := range r.header {
		for _, v := range vs {
			header.Add(k, v)
		}
	}
	if r.code > 0 {
		r.w.WriteHeader(r.code)
	}
	return r.w.Write(r.buffer.Bytes())
}

func (r *ResponseBuffer) Success() bool {
	code := r.StatusCode()
	return code >= http.StatusOK && code < http.StatusBadRequest
}

func (r *ResponseBuffer) Reset() error {
	if r.flushed {
		return errors.New("response flushed")
	}

	r.code = 0
	r.wroteHeader = false
	r.header = http.Header{}
	r.buffer = bytes.Buffer{}
	return nil
}

func (r *ResponseBuffer) StatusCode() int {
	if r.code == 0 {
		return http.StatusOK
	}

	return r.code
}
