package middleware

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"regexp"
	"testing"
)

func TestMethodAndPathSkipper(t *testing.T) {
	type args struct {
		method string
		re     *regexp.Regexp
		r      *http.Request
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"match method and path", args{http.MethodGet, regexp.MustCompile(`/req`), httptest.NewRequest(http.MethodGet, "/req", nil)}, true},
		{"match method only", args{http.MethodGet, regexp.MustCompile(`/req`), httptest.NewRequest(http.MethodGet, "/path", nil)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MethodAndPathSkipper(tt.args.method, tt.args.re)(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MethodAndPathSkipper()() = %v, want %v", got, tt.want)
			}
		})
	}
}
