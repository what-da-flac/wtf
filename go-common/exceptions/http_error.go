package exceptions

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
)

type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewHTTPError(err error) *HTTPError {
	x := &HTTPError{}
	if err != nil {
		x.Message = err.Error()
		st, msg := x.Parse(err)
		if st != nil {
			x.Code = *st
		}
		if msg != nil {
			x.Message = *msg
		}
	}
	return x
}

func (x *HTTPError) WithStatusCode(code int) *HTTPError {
	x.Code = code
	return x
}

func (x *HTTPError) StatusCode() int {
	return x.Code
}

func (x *HTTPError) Error() string {
	return x.Message
}

func (x *HTTPError) WithCustomMessage(message string) *HTTPError {
	x.Message = message
	return x
}

// Parse tries to convert automatically known errors into status codes.
func (x *HTTPError) Parse(err error) (*int, *string) {
	var httpError *HTTPError
	if errors.As(err, &httpError) {
		return &httpError.Code, &httpError.Message
	}
	if errors.Is(err, sql.ErrNoRows) {
		return aws.Int(http.StatusNotFound), aws.String("not found")
	}
	return nil, nil
}
