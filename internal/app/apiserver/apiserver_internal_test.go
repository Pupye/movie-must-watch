package apiserver

import "testing"

import "net/http/httptest"

import "net/http"

import "github.com/stretchr/testify/assert"

func Test_APIServer_HandleHello(t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	s.handleHello().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "fuck yea")
}
