package lib

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
)

type responseBufferTestSuite struct {
	suite.Suite
	recorder *httptest.ResponseRecorder
	buffer   *ResponseBuffer
}

func (r *responseBufferTestSuite) SetupTest() {
	r.recorder = httptest.NewRecorder()
	r.buffer = NewResponseBuffer(r.recorder)
}

func (r *responseBufferTestSuite) TestWriteHeader() {
	// write once
	r.buffer.WriteHeader(http.StatusInternalServerError)
	r.Equal(http.StatusInternalServerError, r.buffer.code)
	r.Equal(http.StatusOK, r.recorder.Code)

	// write again
	r.buffer.WriteHeader(http.StatusNotFound)
	r.Equal(http.StatusInternalServerError, r.buffer.code)
	r.Equal(http.StatusOK, r.recorder.Code)
}

func TestResponseBuffer(t *testing.T) {
	suite.Run(t, &responseBufferTestSuite{})
}
