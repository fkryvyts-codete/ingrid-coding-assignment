// Package http contains logic for handling http requests and translating them to service method calls
package http

import "encoding/json"

type httpError struct {
	message string
	code    int
}

func (e *httpError) StatusCode() int {
	return e.code
}

func (e *httpError) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Message string `json:"message"`
	}{
		Message: e.message,
	})
}

func (e *httpError) Error() string {
	return e.message
}

func newHTTPError(message string, code int) *httpError {
	return &httpError{
		message: message,
		code:    code,
	}
}
